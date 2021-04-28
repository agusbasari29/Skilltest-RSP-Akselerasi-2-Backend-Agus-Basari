package repository

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	InsertTransaction(trx entity.Transaction) (entity.Transaction, error)
	UpdateTransaction(trx entity.Transaction) (entity.Transaction, error)
	GetTransactionByEventAndParticipantAndStatusPayment(trx entity.Transaction) error
	GetTransactionByEventID(trx entity.Transaction) ([]entity.Transaction, error)
	GetTransactionByStatusPayment(trx entity.Transaction) ([]entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) InsertTransaction(trx entity.Transaction) (entity.Transaction, error) {
	err := r.db.Raw("INSERT INTO transactions (participant_id, creator_id, event_id, amount, status_payment, created_at) VALUES (@ParticipantId, @CreatorId, @EventId, @Amount, @StatusPayment, @CreatedAt) WHERE NOT EXIST (SELECT 1 FROM transactions WHERE event_id = @EventId AND participant_id = @ParticipantId)", trx).Save(&trx).Error
	if err != nil {
		return trx, err
	}
	return trx, nil
}

func (r *transactionRepository) UpdateTransaction(trx entity.Transaction) (entity.Transaction, error) {
	err := r.db.Raw("UPDATE transaction SET participant_id = @ParticipantId, creator_id = @CreatorId, event_id = @EventId, amount = @Amount, status_payment = @StatusPayment, updated_at = @UpdatedAt WHERE id = @ID", trx).Save(&trx).Error
	if err != nil {
		return trx, err
	}
	return trx, nil
}

func (r *transactionRepository) GetTransactionByEventAndParticipantAndStatusPayment(trx entity.Transaction) error {
	err := r.db.Raw("SELECT * FROM transaction WHERE participant_id = @ParticipantId AND event_id = @EventId AND status_payment = @StatusPayment", trx).Take(&trx).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) GetTransactionByEventID(trx entity.Transaction) ([]entity.Transaction, error) {
	var trxs []entity.Transaction
	err := r.db.Raw("SELECT * FROM transaction WHERE event_id = @EventId", trx).Find(&trxs).Error
	if err != nil {
		return trxs, err
	}
	return trxs, nil
}

func (r *transactionRepository) GetTransactionByStatusPayment(trx entity.Transaction) ([]entity.Transaction, error) {
	var trxs []entity.Transaction
	err := r.db.Raw("SELECT * FROM transaction WHERE status_payment = @StatusPayment", trx).Find(&trxs).Error
	if err != nil {
		return trxs, err
	}
	return trxs, nil
}
