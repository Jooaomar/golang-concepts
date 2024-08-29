# sync.Mutex

Temos visto como os canais são bons par a comunicação entre coroutines.

Mas e se nós não precisássemos de comunicação? E se nós apenas quiséssemos ter certeza que uma única goroutine pode acessar
uma variável de cada vez para evitar conflitos?

este conceito é chamado de `exclusão mútua`, e o nome convencional para a estrutura de dados que a fornece é `mutex`.

A biblioteca padrão do Go que fornece exclusão mútua com `sync.Mutex` e seus dois métodos:
* Lock
* Unlock

Podemos definir um bloco de código a ser executado em exclusão mútua pelo que o rodeia com uma chamada
para `lock` e `unlock` como mostrado no método `Inc`.

Nós também podemos usar `defer` para garantir que o mutex será desbloqueado como no método `Value`.