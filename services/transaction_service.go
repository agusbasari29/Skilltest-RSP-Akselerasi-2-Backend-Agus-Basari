package services

import (
	"log"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/entity"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/request"
	"github.com/mashingan/smapping"
)

type TransactionServices interface {
	CreateTransaction(req request.RequestTransaction) (entity.Transaction, error)
	GetTransactionByEventAndParticipantAndStatusPayment(trx entity.Transaction) bool
	GetTransactionByEventID(req request.RequestTransaction) ([]entity.Transaction, error)
	GetPendingTrasaction() ([]entity.Transaction, error)
	GetParticipantPendingTrasaction(req request.RequestParticipantTransaction) ([]entity.Transaction, error)
	UpdateTransaction(req request.RequestTransactionUpdate) (entity.Transaction, error)
}

type transactionServices struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionServices(transactionRepository repository.TransactionRepository) *transactionServices {
	return &transactionServices{transactionRepository}
}

func (s *transactionServices) CreateTransaction(req request.RequestTransaction) (entity.Transaction, error) {
	var trx entity.Transaction
	err := smapping.FillStruct(&trx, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	trx.CreatedAt = time.Now()
	createTrx, err := s.transactionRepository.InsertTransaction(trx)
	if err != nil {
		return trx, err
	}
	return createTrx, nil
}

func (s *transactionServices) GetTransactionByEventAndParticipantAndStatusPayment(trx entity.Transaction) bool {
	err := s.transactionRepository.GetTransactionByEventAndParticipantAndStatusPayment(trx)
	return err == nil
}

func (s *transactionServices) GetTransactionByEventID(req request.RequestTransaction) ([]entity.Transaction, error) {
	var trx entity.Transaction
	trx.EventId = int(req.EventId)
	trxs, err := s.transactionRepository.GetTransactionByEventID(trx)
	if err != nil {
		return trxs, err
	}
	return trxs, nil
}

func (s *transactionServices) GetPendingTrasaction() ([]entity.Transaction, error) {
	var trx entity.Transaction
	trx.StatusPayment = ""
	result, err := s.transactionRepository.GetTransactionByStatusPayment(trx)
	if err == nil {
		return result, err
	}
	return result, nil
}

func (s *transactionServices) GetParticipantPendingTrasaction(req request.RequestParticipantTransaction) ([]entity.Transaction, error) {
	var trx entity.Transaction
	trx.StatusPayment = ""
	trx.ParticipantId = int(req.ParticipantId)
	result, err := s.transactionRepository.GetTransactionByStatusPayment(trx)
	if err == nil {
		return result, err
	}
	return result, nil
}

func (s *transactionServices) UpdateTransaction(req request.RequestTransactionUpdate) (entity.Transaction, error) {
	var trx entity.Transaction
	err := smapping.FillStruct(&trx, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	trx.UpdatedAt = time.Now()
	update, err := s.transactionRepository.UpdateTransaction(trx)
	if err != nil {
		return trx, err
	}
	return update, nil
}
