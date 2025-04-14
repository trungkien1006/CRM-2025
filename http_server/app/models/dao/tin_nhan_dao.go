package dao

import (
	"admin-v1/app/helpers"
	"admin-v1/app/models/db"
	"admin-v1/app/models/requests"
	"errors"
)

func CreateMessageExec(req *requests.Tin_nhan_create) error {
	tx := helpers.GormDB.Begin()

	var count int64

	if err := tx.Debug().
		Table("nhan_vien").
		Where("id = ?", req.Sender_id).
		Count(&count).Error;
	err != nil{
		tx.Rollback()
		return errors.New("co loi khi kiem tra nhan vien gui tin nhan ton tai: "+ err.Error())
	}

	if count == 0 {
		tx.Rollback()
		return errors.New("nhan vien gui tin nhan khong ton tai")
	}

	count = 0

	if err := tx.Debug().
		Table("nhan_vien").
		Where("id = ?", req.Receiver_id).
		Count(&count).Error;
	err != nil{
		tx.Rollback()
		return errors.New("co loi khi kiem tra nhan vien nhan tin nhan ton tai: "+ err.Error())
	}

	var message db.Tin_nhan = db.Tin_nhan{
		Sender_id: req.Sender_id,
		Receiver_id: req.Receiver_id,
		Content: req.Content,
		Is_read: false,
	}

	if err := tx.Debug().Create(&message).Error; err != nil {
		tx.Rollback()
		return errors.New("co loi khi them tin nhan: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}

func CreateBatchMessageExec(req *requests.Tin_nhan_create_batch) error {
	tx := helpers.GormDB.Begin()

	var count 			int64

	senderIdMap 	:= make(map[int] int)
	receiverIdMap 	:= make(map[int] int)

	for _, value := range req.Tin_nhan_batch {
		senderIdMap[value.Sender_id] 		= value.Sender_id
		receiverIdMap[value.Receiver_id] 	= value.Receiver_id
	}

	senderIdArr 	:= make([]int, 0, len(senderIdMap))
	receiverIdArr 	:= make([]int, 0, len(receiverIdMap))

	for _, id := range senderIdMap {
		senderIdArr = append(senderIdArr, id)
	}

	for _, id := range receiverIdMap {
		receiverIdArr = append(receiverIdArr, id)
	}

	if err := tx.Debug().
		Table("nhan_vien").
		Where("id IN ?", senderIdArr).
		Count(&count).Error;
	err != nil{
		tx.Rollback()
		return errors.New("co loi khi kiem tra nhan vien gui tin nhan ton tai: "+ err.Error())
	}

	if count < int64(len(senderIdArr)) {
		tx.Rollback()
		return errors.New("1 so nhan vien gui tin nhan khong ton tai")
	}

	count = 0

	if err := tx.Debug().
		Table("nhan_vien").
		Where("id IN ?", receiverIdArr).
		Count(&count).Error;
	err != nil{
		tx.Rollback()
		return errors.New("co loi khi kiem tra nhan vien nhan tin nhan ton tai: "+ err.Error())
	}

	if count < int64(len(receiverIdArr)) {
		tx.Rollback()
		return errors.New("1 so nhan vien nhan tin nhan khong ton tai")
	}

	messageBatch := make([]db.Tin_nhan, 0, len(req.Tin_nhan_batch))

	for _, value := range req.Tin_nhan_batch {
		messageBatch = append(messageBatch, db.Tin_nhan{
			Sender_id: value.Sender_id,
			Receiver_id: value.Receiver_id,
			Content: value.Content,
			Is_read: false,
		})
	}

	if err := tx.Debug().Create(&messageBatch).Error; err != nil {
		tx.Rollback()
		return errors.New("co loi khi them tin nhan: " + err.Error())
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("loi commit transaction: " + err.Error())
	}

	return nil
}