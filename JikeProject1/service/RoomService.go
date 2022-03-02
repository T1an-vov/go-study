package service

import (
	"JikeProject1/dao"
	"JikeProject1/module"
)

func RoomCreate(room *module.Room) error {
	if err:=dao.DB.Save(room).Error;err!=nil{
		return err
	}
	return nil
}

func RoomDelete(number string,owner string)(bool,error){
	room:=new(module.Room)
	if err:=dao.DB.Where("number= ?",number).First(room).Error;err!=nil{
		return true,err
	}
	if room.Owner!=owner{
		return false,nil
	}
	dao.DB.Delete(room)
	return true,nil
}

func RoomCheck(number string) *module.Room {
	room:=new(module.Room)
	if err:=dao.DB.Where("number=?",number).First(room).Error;err!=nil{
		return nil
	}
	return room
}
func RoomChangeName(number string,newName string)error  {
	room:=new(module.Room)
	if err:=dao.DB.Model(room).Where("number=?",number).Update("name",newName).Error;err!=nil{
		return err
	}
	return nil
}

func RoomChangeMax(number string,newMax int)error  {
	room:=new(module.Room)
	if err:=dao.DB.Model(room).Where("number=?",number).Update("max",newMax).Error;err!=nil{
		return err
	}
	return nil
}

func RoomChangeOwner(number string,newOwner string)error  {
	room:=new(module.Room)
	if err:=dao.DB.Model(room).Where("number=?",number).Update("owner",newOwner).Error;err!=nil{
		return err
	}
	return nil
}

func RoomsGet(name string) *[]module.Room {
	rooms:=new([]module.Room)
	dao.DB.Where("name like ?","%"+name+"%").Order("create_at").Find(rooms)
	return rooms
}