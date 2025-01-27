# Multicall2 - Consulta de Saldos WETH

Este projeto implementa uma API REST que utiliza a biblioteca multicall da Omnes Tech para consultar saldos WETH de múltiplos endereços Ethereum em uma única chamada.

## Importância do Multicall

A biblioteca multicall da Omnes Tech oferece várias vantagens importantes:

1. **Eficiência**: Permite agrupar múltiplas chamadas de leitura em uma única transação, reduzindo significativamente o número de requisições RPC.
2. **Economia**: Reduz custos de infraestrutura e latência ao minimizar o número de chamadas ao nó Ethereum.
3. **Atomicidade**: Garante que todos os dados sejam lidos no mesmo bloco, oferecendo uma visão consistente do estado da blockchain.
4. **Performance**: Ideal para aplicações que precisam consultar múltiplos contratos ou dados simultaneamente.

## Instalação

```bash
# Clone o repositório
git clone https://github.com/Afonsodalvi/multicall-readBalance.git

# Entre no diretório
cd multicall-readBalance

# Instale as dependências
go mod tidy
```

## Rotas da API

### GET /health
Verifica se o serviço está funcionando.

**Resposta**:
```json
{
    "status": "up"
}
```

### POST /getBalance
Consulta os saldos WETH de múltiplos endereços.

**Payload**:
```json
{
    "addresses": [
        "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
        "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
    ]
}
```

**Resposta de Sucesso**:
```json
{
    "status": "success",
    "balances": [
        {
            "address": "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
            "balance": "1000000000000000000"
        },
        {
            "address": "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640",
            "balance": "2000000000000000000"
        }
    ]
}
```

## Exemplos de Uso com cURL

1. **Consulta Simples**:
```bash
curl -X POST http://localhost:8080/getBalance \
-H "Content-Type: application/json" \
-d '{
    "addresses": ["0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"]
}'
```

2. **Consulta Múltipla**:
```bash
curl -X POST http://localhost:8080/getBalance \
-H "Content-Type: application/json" \
-d '{
    "addresses": [
        "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",
        "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640",
        "0x6B175474E89094C44Da98b954EedeAC495271d0F"
    ]
}'
```

## Observações Importantes

1. Os saldos são retornados em wei (18 casas decimais)
2. Para converter para WETH, divida o valor por 10^18
3. Endereços inválidos resultarão em erro
4. O serviço usa a rede Ethereum mainnet
5. O contrato WETH utilizado é: `0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2`

## Tratamento de Erros

A API retorna os seguintes códigos de erro:

- `400`: Request inválido (endereços mal formatados ou payload incorreto)
- `405`: Método não permitido (apenas POST é aceito)
- `500`: Erro interno do servidor

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## Suporte

Para suporte, abra uma issue no repositório ou entre em contato com a equipe de desenvolvimento.
```

Esta documentação fornece uma visão completa do projeto, incluindo:
1. Explicação da importância do multicall
2. Instruções de instalação
3. Documentação detalhada das rotas
4. Exemplos práticos de uso
5. Informações sobre tratamento de erros
6. Guia de contribuição

Você pode personalizar ainda mais adicionando:
- Exemplos específicos do seu caso de uso
- Mais detalhes sobre a configuração
- Informações sobre deployment
- Links para documentação adicional
