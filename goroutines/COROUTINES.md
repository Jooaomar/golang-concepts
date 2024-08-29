# Goroutines

Uma goroutine é um segmento leve e gerenciado pelo runtime de Go.

```
go f(x, y, z)
```

inicia uma nova execução goroutine

```
f(x, y, z)
```

A avaliação de `f, x, y, e z` acontece na goroutine corrente e para a execução de f acontece em uma goroutine nova.

Goroutines executam no mesmo espaço de endereço, para que o acesso à memória compartilhada seja sincronizada. O pacote sync fornece as primitivas úteis, embora você não vá precisar muito deles em Go, pois há outras primitivas. (Veja o próximo slide).