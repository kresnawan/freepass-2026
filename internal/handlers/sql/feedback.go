package sql

import (
	"canteen/internal/models"
	"canteen/internal/storage/mariadb"
)

func SelectCanteenFeedback(cid string) ([]models.Feedback, error) {
	var feedbacks = make([]models.Feedback, 0)
	query := `
		SELECT
			*
		FROM
			feedback f
		JOIN
			order o
		ON
			o.order_id = f.order_id
		WHERE
			o.canteen_id = ?

	`
	res, err := mariadb.Db.Query(query, cid)
	if err != nil {
		return feedbacks, err
	}

	for res.Next() {
		var feedback models.Feedback
		err := res.Scan(
			&feedback.FeedbackId,
			&feedback.OrderId,
			&feedback.Description,
			&feedback.Star,
		)

		if err != nil {
			return feedbacks, err
		}
	}

	return feedbacks, nil
}
