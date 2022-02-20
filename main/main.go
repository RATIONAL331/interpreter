package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	curUser, e := user.Current()
	if e != nil {
		panic(any(e))
	}

	/**
	let map = fn(arr, f) {
		let iter = fn(arr, accumulated) {
			if (len(arr) == 0) {
				accumulated
			} else {
				iter(rest(arr), push(accumulated, f(first(arr))));
			}
		};
		iter(arr, []);
	};
	let a = [1, 2, 3, 4];
	let double = fn(x) { x * 2 };
	map(a, double);

	let map = fn(arr, f) { let iter = fn(arr, accumulated) { if (len(arr) == 0) {accumulated} else {iter(rest(arr), push(accumulated, f(first(arr))));}}; iter(arr, []);}; let a = [1, 2, 3, 4]; let double = fn(x) { x * 2 }; map(a, double);
	*/

	/**
	let reduce = fn(arr, initial, f) {
		let iter = fn(arr, result) {
			if(len(arr) == 0) {
				result
			} else {
				iter(rest(arr), f(result, first(arr)));
			}
		};
		iter(arr, initial);
	};
	let sum = fn(arr) {
		reduce(arr, 0, fn(initial, el) { initial + el });
	}
	sum([1, 2, 3, 4, 5]);

	 let reduce = fn(arr, initial, f) {let iter = fn(arr, result) {if(len(arr) == 0) {result} else {iter(rest(arr), f(result, first(arr)));}};iter(arr, initial);}; let sum = fn(arr) {reduce(arr, 0, fn(initial, el) { initial + el });} sum([1, 2, 3, 4, 5]);
	*/

	/**
	let unless = macro(condition, consequence, alternative) {
		quote(if (!(unquote(condition))){
				unquote(consequence);
				}
				else { unquote(alternative);
				}
		);
	};
	unless(10 > 5, puts("not greater"), puts("greater"));

	let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); }; unless(10 > 5, puts("not greater"), puts("greater"));
	*/

	fmt.Printf("Hello %s! This is the Interpreter programming lanuguage!\n", curUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
