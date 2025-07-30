CREATE DATABASE filmes_criticas;

USE filmes_criticas;

-- Tabela de usuários
CREATE TABLE usuario (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) DEFAULT (UUID()),
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    Nomeuser VARCHAR(255) NOT NULL UNIQUE,
    senha VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

INSERT INTO usuario (nome, email, Nomeuser, senha)
VALUES ('Sara Lance', 'sara.lance@canario.com', 'WhiteCanary', 'senhaSuperSegura123');

select * from usuario

-- Tabela de filmes
CREATE TABLE filme (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) DEFAULT (UUID()),
    titulo VARCHAR(255) NOT NULL,
    sinopse TEXT NOT NULL,
    diretor VARCHAR(255),
    data_lancamento DATE,
    poster_url VARCHAR(255),
    genero VARCHAR(55),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- Tabela de avaliações
CREATE TABLE avaliacoes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    usuario_id INT,
    filme_id INT,
    nota INT CHECK (nota BETWEEN 1 AND 5),
    comentario TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (usuario_id) REFERENCES usuario(id),
    FOREIGN KEY (filme_id) REFERENCES filme(id)
);
describe filme
ALTER TABLE Filme ADD COLUMN codigo_externo INT NULL;
