package main

import (
	"context"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"
)

func handler(ctx context.Context, event events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	body := `
	<!DOCTYPE html>
	<html lang="zh">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>猜数字小游戏</title>
		<style>
			div, ul, li, h1 {
				margin: 0;
				padding: 0;
				border: 0;
				font-size: 100%;
				font: inherit;
				vertical-align: baseline;
			}
			.game-title {
				font-size: 32px;
				font-weight: bold;
			}
			#number-input {
				width: 50px;
			}
			.result-list-item {
				list-style-type: none;
				color: black;
			}
			.result-list-item-wrong {
				background-color: #d85140;
			}
			.result-list-item-right {
				background-color: #58a55c;
			}
		</style>
	</head>
	<body>
		<div>
			<div>
				<h1 class='game-title'>猜数字小游戏</h1>
				<div>我刚才随机选定了一个1~100之间的整数，试试看你多少次能够猜中它。每次我都会告诉你猜的数字是大了还是小了。</div>
				<input id="number-input" type="number" min="1" max="100" step="1">
				<button class="guess-button" type="button">我猜</button>
				<button class="reset-button" type="button">重新开始</button>
			</div>
			<ul class="result-list"></ul>
		</div>
		<script type="text/javascript">
		window.gameContext = {};
	
		function resetGame() {
			gameContext.number = Math.floor(Math.random() * 100) + 1;
			gameContext.guessRound = 1;
			gameContext.finished = false;
			console.log(gameContext);
		}
	
		function appendResult(guessRound, guessNumber, rightNumber) {
			console.log(rightNumber);
			let listItem = document.createElement('li');
			if (guessNumber < rightNumber) {
				listItem.innerText = '你第' + guessRound + '次猜的数字是:' + guessNumber + '，小了！';
				listItem.className = 'result-list-item result-list-item-wrong';
			}
			else if (guessNumber > rightNumber) {
				listItem.innerText = '你第' + guessRound + '次猜的数字是:' + guessNumber + '，大了！';
				listItem.className = 'result-list-item result-list-item-wrong';
			}
			else {
				listItem.innerText = '你第' + guessRound + '次猜的数字是:' + guessNumber + '，正确！';
				listItem.className = 'result-list-item result-list-item-right';
			}
			
			document.querySelector('.result-list').appendChild(listItem);
		}
	
		function clearResultList() {
			let resultList = document.querySelector('.result-list');
			while (resultList.firstChild) {
				resultList.removeChild(resultList.firstChild);
			}
		}
	
		function resetButtonClick() {
			resetGame();
			clearResultList();
		}
	
		function updateGameContext(guessNumber) {
			if (gameContext.finished) {
				return;
			}
			gameContext.guessRound++;
			if (gameContext.number == guessNumber) {
				gameContext.finished = true;
			}
		}
	
		function guessButtonClick() {
			if (gameContext.finished) {
				return;
			}
			let guessNumber = document.querySelector('#number-input').value;
			appendResult(gameContext.guessRound, guessNumber, gameContext.number);
			updateGameContext(guessNumber);
		}
	
		function bindEvents() {
			document.querySelector('.reset-button').addEventListener('click', resetButtonClick);
			document.querySelector('.guess-button').addEventListener('click', guessButtonClick);
		}
	
		window.onload = function() {
			resetGame();
			bindEvents();
		};
	
		</script>
	</body>
	</html>
    `
	resp := events.APIGatewayResponse{
		IsBase64Encoded: false,
		Headers:         map[string]string{"Content-Type": "text/html"},
		StatusCode:      200,
		Body:            body,
	}
	return resp, nil
}

func main() {
	cloudfunction.Start(handler)
}
