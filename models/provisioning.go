package models

import "github.com.br/rodrygo262/hackaton_golang/db"

type Provisioning struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	OrderId  string `gorm:"size:255;not null;unique" json:"order_id"`
	ClientId string `gorm:"size:255;not null;unique" json:"client_id"`
	SKU      string `gorm:"size:255;not null;unique" json:"sku"`
}

func (prov *Provisioning) Create_prov() (*Provisioning, error) {

	conexao := db.Connect()
	defer conexao.Close()

	var err error
	err = conexao.Debug().Model(&Provisioning{}).Create(&prov).Error
	if err != nil {
		return &Provisioning{}, err
	}

	return prov, nil
}

func (prov *Provisioning) List_all_provisiongs() ([]Provisioning, error) {

	conexao := db.Connect()
	defer conexao.Close()

	provisionigns := []Provisioning{}
	var err error
	err = conexao.Debug().Model(&Provisioning{}).Limit(100).Find(&provisionigns).Error
	if err != nil {
		return []Provisioning{}, err
	}

	return provisionigns, nil
}
