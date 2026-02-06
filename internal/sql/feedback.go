package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
	"database/sql"
	"errors"
	"strconv"

	"github.com/oklog/ulid/v2"
)

type PublicFeedback struct {
	FeedbackId       int    `json:"feedback_id"`
	CanteenName      string `json:"canteen_name"`
	CustomerUsername string `json:"customer_username"`
	Description      string `json:"description"`
	Star             int    `json:"star"`
}

func SelectCanteenFeedback(cid string, page string) ([]PublicFeedback, error) {
	var feedbacks = make([]PublicFeedback, 0)

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return feedbacks, err
	}

	query := `
		SELECT
			f.feedback_id,
			c.name AS canteen_name,
			acc.username AS customer_username,
			f.description,
			f.star
		FROM
			feedback f
		JOIN
			` + "`order`" + ` o ON o.order_id = f.order_id
		JOIN
			canteen c ON c.canteen_id = o.canteen_id
		JOIN
			accounts acc ON acc.account_id = o.customer_id
		WHERE
			o.canteen_id = ?
		LIMIT
			10
		OFFSET
			?
	`
	res, err := mariadb.Db.Query(query, cid, ((pageInt - 1) * 10))
	if err != nil {
		return feedbacks, err
	}

	for res.Next() {
		var feedback PublicFeedback
		err := res.Scan(
			&feedback.FeedbackId,
			&feedback.CanteenName,
			&feedback.CustomerUsername,
			&feedback.Description,
			&feedback.Star,
		)

		if err != nil {
			return feedbacks, err
		}

		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}

func AddUserFeedback(oid ulid.ULID, desc string, star int) error {
	query := `
		SELECT
			status
		FROM
			` + "`order`" + `
		WHERE
			order_id = ?
	`

	var status int
	err := mariadb.Db.QueryRow(query, oid).Scan(&status)

	if err != nil {
		return err
	}

	if status != 9 {
		return errors.New("This order did not completed yet, feedback did not allowed")
	}

	err = InsertFeedback(oid, desc, star)
	if err != nil {
		return err
	}

	return nil
}

func InsertFeedback(oid ulid.ULID, desc string, star int) error {
	query := `
		INSERT INTO
			feedback
			(order_id, description, star)
		VALUES
			(?, ?, ?)
	`

	_, err := mariadb.Db.Exec(query, oid, desc, star)
	if err != nil {
		return err
	}

	return nil
}

func GetMyOrderFeedback(cusid ulid.ULID) ([]models.Feedback, error) {
	var feedbacks = make([]models.Feedback, 0)
	query := `
		SELECT
			f.feedback_id,
			f.order_id,
			o.canteen_id,
			o.customer_id,
			f.description,
			f.star
		FROM
			feedback f
		JOIN
			` + "`order`" + ` o
		ON
			o.order_id = f.order_id
		WHERE
			o.customer_id = ?

	`
	res, err := mariadb.Db.Query(query, cusid)
	if err != nil {
		return feedbacks, err
	}

	for res.Next() {
		var feedback models.Feedback
		err := res.Scan(
			&feedback.FeedbackId,
			&feedback.OrderId,
			&feedback.CanteenId,
			&feedback.CustomerId,
			&feedback.Description,
			&feedback.Star,
		)

		if err != nil {
			return feedbacks, err
		}
	}

	return feedbacks, nil
}

func DeleteFeedback(fid string) error {
	query := `
		DELETE FROM
			feedback
		WHERE
			feedback_id = ?
	`

	_, err := mariadb.Db.Exec(query, fid)
	if err != nil {
		return err
	}

	return nil
}

func CheckFeedbackOwnership(fid int, uid ulid.ULID) (bool, error) {
	query := `
		SELECT
			o.order_id
		FROM
			feedback f
		JOIN
			` + "`order`" + ` o ON f.order_id = o.order_id
		JOIN
			canteen_ownership co ON co.canteen_id = o.canteen_id
		WHERE
			f.feedback_id = ? AND co.owner_id = ?
	`
	var canteenId ulid.ULID
	err := mariadb.Db.QueryRow(query, fid, uid).Scan(&canteenId)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
