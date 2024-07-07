# API Leis de Ohm em Go

Este projeto consiste em uma API simples em Go para realizar cálculos elétricos básicos utilizando as Leis de Ohm. A API oferece endpoints para calcular tensão, corrente, resistência, resistividade, energia e potência com base nos dados fornecidos.

Fiz essa API em base a uma lista de exercicios do meu professor da materia de Eletroeltronica. Professor Djalma. 

## Endpoints Disponíveis

- `/CalcularTensao`: Calcula a Tensão (V) usando a Lei de Ohm (V = I * R).
- `/CalcularCorrente`: Calcula a Corrente (I) usando a Lei de Ohm (I = V / R).
- `/CalcularResistencia`: Calcula a Resistência (R) usando a Lei de Ohm (R = V / I).
- `/CalcularResistividade`: Calcula a Resistividade (ρ = (R * A) / L).
- `/CalcularEnergia`: Calcula a Energia Elétrica (E = P * t).
- `/CalcularEnergiaVIT`: Calcula a Energia usando Tensão, Corrente e Tempo (E = V * I * t).
- `/CalcularEnergiaCRT`: Calcula a Energia usando Corrente, Resistência e Tempo (E = I^2 * R * t).
- `/CalcularEnergiaVRT`: Calcula a Energia usando Tensão, Resistência e Tempo (E = (V^2) / R * t).
- `/CalcularPotenciaVI`: Calcula a Potência usando Tensão e Corrente (P = V * I).
- `/CalcularPotenciaIR`: Calcula a Potência usando Corrente e Resistência (P = I^2 * R).
- `/CalcularPotenciaVR`: Calcula a Potência usando Corrente ao quadrado e Resistência (P = (I^2) * R).

## Funcionamento

A API espera um corpo de requisição no formato JSON contendo os dados necessários para realizar os cálculos solicitados. Os resultados são retornados como JSON com a resposta correspondente.

## Uso

Para utilizar a API, envie requisições HTTP POST para os endpoints listados acima, especificando os dados necessários no corpo da requisição em formato JSON.

Exemplo de requisição para calcular a Tensão:

  http://localhost:8080/CalcularTensao 
  body > raw json format 
  {
  "corrente": 5,
  "resistencia": 10
  
}'

No arquivo main esta especificado como chamar cada variavel para fazer o calculo. Ficarei grato para dicas e melhorias para meu codigo.
