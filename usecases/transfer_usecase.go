package usecases

import (
	"errors"
	"time"
	"workshop4-backend/domain/entities"
	"workshop4-backend/domain/repositories"
)

type TransferUseCase struct {
	transferRepo repositories.TransferRepository
	userRepo     repositories.UserRepository
}

func NewTransferUseCase(transferRepo repositories.TransferRepository, userRepo repositories.UserRepository) *TransferUseCase {
	return &TransferUseCase{
		transferRepo: transferRepo,
		userRepo:     userRepo,
	}
}

func (uc *TransferUseCase) TransferPoints(fromUserID uint, req *entities.TransferRequest) (*entities.Transfer, error) {
	// Validate input
	if fromUserID == 0 {
		return nil, errors.New("invalid from user ID")
	}

	if fromUserID == req.ToUserID {
		return nil, errors.New("cannot transfer points to yourself")
	}

	// Check if from user exists and has enough points
	fromUser, err := uc.userRepo.GetByID(fromUserID)
	if err != nil {
		return nil, errors.New("from user not found")
	}

	if fromUser.Points < req.Amount {
		return nil, errors.New("insufficient points")
	}

	// Check if to user exists
	toUser, err := uc.userRepo.GetByID(req.ToUserID)
	if err != nil {
		return nil, errors.New("to user not found")
	}

	// Begin transaction
	tx := uc.transferRepo.BeginTransaction()

	// Create transfer record with current timestamp
	now := time.Now()
	transfer := &entities.Transfer{
		FromUserID:    fromUserID,
		ToUserID:      req.ToUserID,
		Amount:        req.Amount,
		Description:   req.Description,
		Status:        "completed",
		TransferredAt: now,
	}

	err = uc.transferRepo.CreateWithTransaction(tx, transfer)
	if err != nil {
		uc.transferRepo.RollbackTransaction(tx)
		return nil, errors.New("failed to create transfer record")
	}

	// Update from user points (deduct)
	err = uc.userRepo.UpdatePointsWithTransaction(tx, fromUserID, fromUser.Points-req.Amount)
	if err != nil {
		uc.transferRepo.RollbackTransaction(tx)
		return nil, errors.New("failed to deduct points from sender")
	}

	// Update to user points (add)
	err = uc.userRepo.UpdatePointsWithTransaction(tx, req.ToUserID, toUser.Points+req.Amount)
	if err != nil {
		uc.transferRepo.RollbackTransaction(tx)
		return nil, errors.New("failed to add points to receiver")
	}

	// Commit transaction
	err = uc.transferRepo.CommitTransaction(tx)
	if err != nil {
		uc.transferRepo.RollbackTransaction(tx)
		return nil, errors.New("failed to complete transfer")
	}

	return transfer, nil
}

func (uc *TransferUseCase) GetTransferHistory(userID uint, limit, offset int) (*entities.TransferHistoryResponse, error) {
	if userID == 0 {
		return nil, errors.New("invalid user ID")
	}

	// Check if user exists
	_, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	transfers, total, err := uc.transferRepo.GetTransferHistory(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	// Convert to response format
	transferResponses := make([]entities.TransferResponse, len(transfers))
	for i, transfer := range transfers {
		transferResponses[i] = entities.TransferResponse{
			ID:            transfer.ID,
			FromUserID:    transfer.FromUserID,
			ToUserID:      transfer.ToUserID,
			Amount:        transfer.Amount,
			Description:   transfer.Description,
			Status:        transfer.Status,
			TransferredAt: transfer.TransferredAt,
			CreatedAt:     transfer.CreatedAt,
		}
	}

	return &entities.TransferHistoryResponse{
		Transfers: transferResponses,
		Total:     total,
	}, nil
}

func (uc *TransferUseCase) GetTransferByID(transferID uint) (*entities.Transfer, error) {
	if transferID == 0 {
		return nil, errors.New("invalid transfer ID")
	}

	transfer, err := uc.transferRepo.GetByID(transferID)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}
