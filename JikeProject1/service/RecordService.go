package service

import (
	"JikeProject1/dao"
	"JikeProject1/module"
	"errors"
)

func RecordCreate(email string,number string)error  {
	room:=new(module.Room)
	record:=new(module.Record)
	if err:=dao.DB.Where("number=?",number).First(room).Error;err!=nil{
		return errors.New("该房间不存在")
	}
	if err:=dao.DB.Where("email=? and number = ? ",email,number).First(record).Error;err==nil{
		return errors.New("已加入该房间")
	}
	if room.Access=="no"{
		return errors.New("该房间不允许加入")
	}
	if room.People==room.Max{
		return errors.New("该房间人数已满")
	}
	record.Email=email
	record.Number=number
	room.People+=1
	dao.DB.Save(record)
	dao.DB.Save(room)
	return nil
}

func RecordDelete(email string,number string)error  {
	record:=new(module.Record)
	if err:=dao.DB.Where("email = ? and number = ? ",email,number).First(record).Error;err!=nil{
		return errors.New("还未加入该房间")
	}
	dao.DB.Delete(record)
	return nil
}


func RecordGet(number string) *[]module.Record {
	records:=new([]module.Record)
	dao.DB.Debug().Where("number=?",number).Find(records)
	return records
}


//房主删除房间
func RecordDeleteRoom(email string,number string)error{
	room:=new(module.Room)
	if err:=dao.DB.Where("number=?",number).First(room).Error;err!=nil{
		return errors.New("该房间不存在")
	}
	if room.Owner!=email{
		return errors.New("只有房主可以删除房间")
	}
	dao.DB.Where("number = ?",number).Delete(module.Record{})
	dao.DB.Delete(room)
	return nil
}


//房主删除用户
func RecordDeleteUser(email string,number string,userEmail string)error  {
	room:=new(module.Room)
	record:=new(module.Record)
	if err:=dao.DB.Where("number=?",number).First(room).Error;err!=nil{
		return errors.New("该房间不存在")
	}
	if room.Owner!=email{
		return errors.New("只有房主可以删除用户")
	}
	if err:=dao.DB.Where("number=? and email =?",number,userEmail).First(record).Error;err!=nil{
		return errors.New("房间里没有该用户")
	}
	room.People-=1
	dao.DB.Save(room)
	dao.DB.Delete(record)
	return nil
}