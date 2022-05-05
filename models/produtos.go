package models

import "gabriel/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBanco()

	selectTodosOsProdutos, err := db.Query("select * from produtoslojas")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	insereDadosNoBanco, err := db.Prepare("insert into produtoslojas(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func DeletaProduto(id string) {
	db := db.ConectaBanco()
	deletarOProduto, err := db.Prepare("delete from produtoslojas where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close()
}
func EditaProduto(id string) Produto {
	db := db.ConectaBanco()

	produtoDoBanco, err := db.Query("select from produtoslojas where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}
	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}
func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	AtualizaProduto, err := db.Prepare("update produtoslojas set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
