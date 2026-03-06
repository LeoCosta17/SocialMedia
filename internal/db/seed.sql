INSERT INTO users (email, username, password, created_at) 
VALUES 
('joao.silva@email.com', 'joaosilva', 'senha_segura_123', CURRENT_TIMESTAMP),
('maria.oliveira@email.com', 'maria_oliveira', 'minha_senha_456', CURRENT_TIMESTAMP),
('tech.user@provider.com', 'tech_user', 'access_2026', '2026-02-27 10:00:00-0300'),
('suporte@empresa.com', 'admin_sup', 'root_pass_789', CURRENT_TIMESTAMP),
('contato@site.com', 'contato_geral', 'contato_99', CURRENT_TIMESTAMP);

INSERT INTO posts (title, user_id, content, tags, created_at, updated_at) 
VALUES 
('Dicas de SQL', 1, 'Conteúdo sobre boas práticas em bancos de dados.', '{"SQL", "Educação"}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Meu Segundo Post', 2, 'Este é um post de exemplo para teste de relacionamento.', '{"Teste"}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Configurando o Docker', 3, 'Passo a passo para subir um container com Postgres.', '{"Docker", "DevOps"}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Explorando JSONB', 4, 'Como trabalhar com tipos complexos no PostgreSQL.', '{"Avançado", "JSON"}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Performance em Queries', 5, 'Entenda como o EXPLAIN ANALYZE funciona.', '{"Performance", "SQL"}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Novidades do Framework', 6, 'Um resumo das últimas atualizações da semana.', '{"News", "Tech"}', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO comments (post_id, user_id, content, created_at) 
VALUES 
-- Comentários para o Post 1
(1, 2, 'Excelente explicação!', CURRENT_TIMESTAMP),
(1, 3, 'Me ajudou muito no projeto da faculdade.', CURRENT_TIMESTAMP),
(1, 4, 'Pode falar mais sobre a instalação?', CURRENT_TIMESTAMP),

-- Comentários para o Post 2
(2, 1, 'Concordo plenamente com o que foi dito.', CURRENT_TIMESTAMP),
(2, 5, 'Interessante, não tinha pensado por esse lado.', CURRENT_TIMESTAMP),

-- Comentários para o Post 3
(3, 6, 'Docker facilita demais a vida, ótimo post!', CURRENT_TIMESTAMP),
(3, 2, 'Tive um erro na porta 5432, pode ajudar?', CURRENT_TIMESTAMP),
(3, 4, 'Ficou muito claro o passo a passo.', CURRENT_TIMESTAMP),
(3, 1, 'Top!', CURRENT_TIMESTAMP),

-- Comentários para o Post 4
(4, 3, 'JSONB é vida no Postgres!', CURRENT_TIMESTAMP),
(4, 5, 'Quais as vantagens em relação ao JSON comum?', CURRENT_TIMESTAMP),

-- Comentários para o Post 5
(5, 6, 'Sempre uso o EXPLAIN, essencial.', CURRENT_TIMESTAMP),
(5, 2, 'Dica de ouro para quem quer performance.', CURRENT_TIMESTAMP),
(5, 4, 'Faz um post sobre índices depois!', CURRENT_TIMESTAMP),

-- Comentários para o Post 6
(6, 1, 'As novidades dessa semana estão incríveis.', CURRENT_TIMESTAMP),
(6, 3, 'Onde encontro o link da documentação oficial?', CURRENT_TIMESTAMP),

-- Comentários para o Post 7
(7, 5, 'Muito bom o conteúdo!', CURRENT_TIMESTAMP),
(7, 6, 'Acompanhando sempre.', CURRENT_TIMESTAMP),
(7, 2, 'Valeu por compartilhar.', CURRENT_TIMESTAMP),
(7, 4, 'Show de bola.', CURRENT_TIMESTAMP),

-- Comentários para o Post 8
(8, 1, 'Esperando a parte 2!', CURRENT_TIMESTAMP),
(8, 3, 'Sensacional o texto.', CURRENT_TIMESTAMP),
(8, 5, 'Aprendi muito hoje.', CURRENT_TIMESTAMP);