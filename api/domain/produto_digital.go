package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type ProdutoDigial struct {
	ID              string    `json:"id" gorm:"type:uuid;primary_key"`
	ProdutoID       string    `json:"-" gorm:"column:produto_id;type:uuid;notnull"`
	Produto         Produto   `json:"produto" valid:"notnull"`
	Link            string    `json:"link"`
	DataCriacao     time.Time `json:"data_criacao" valid:"-"`
	DataModificacao time.Time `json:"data_modificacao" valid:"-"`
}

func NewProdutoDigital(produto Produto) (*ProdutoDigial, error) {
	pd := ProdutoDigial{
		Produto: produto,
	}
	pd.prepare()
	err := pd.Validate()

	if err != nil {
		return nil, err
	}

	return &pd, nil
}

func (pd *ProdutoDigial) prepare() {
	pd.ID = uuid.NewV4().String()
	pd.DataCriacao = time.Now()
	pd.DataModificacao = time.Now()
}

func (pd ProdutoDigial) Validate() error {
	_, err := govalidator.ValidateStruct(pd)
	return err
}