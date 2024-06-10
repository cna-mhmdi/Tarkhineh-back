package db

import "context"

type CreateProfileTxParams struct {
	CreateProfileParams
	AfterCreate func(profile Profile) error
}

type CreateProfileTxResult struct {
	Profile Profile
}

func (store *SQLStore) CreateProfileTx(ctx context.Context, arg CreateProfileTxParams) (CreateProfileTxResult, error) {
	var result CreateProfileTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Profile, err = q.CreateProfile(ctx, arg.CreateProfileParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.Profile)
	})

	return result, err
}
