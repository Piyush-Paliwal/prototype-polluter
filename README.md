# Prototype-Polluter

Take a list of domains and check for Prototype Pollution Vulnerability.

## Install

```
▶ go get github.com/PiyushThePal/prototype-polluter
```
## Basic Usage

Use "-h" for help

```
▶ cat domains-list.txt | prototype-polluter
```
<img width="1018" alt="Screenshot 2021-07-10 at 1 40 35 PM" src="https://user-images.githubusercontent.com/46394419/125156722-93690c80-e184-11eb-97d1-aac6409afda0.png">

```
▶ cat domains-list.txt | prototype-polluter -v
```
<img width="1017" alt="Screenshot 2021-07-10 at 1 34 10 PM" src="https://user-images.githubusercontent.com/46394419/125156712-864c1d80-e184-11eb-92f1-4202be0bf661.png">

```
▶ waybackurls example.com | prototype-polluter -v
```
<img width="799" alt="Screenshot 2021-07-10 at 1 33 25 PM" src="https://user-images.githubusercontent.com/46394419/125156723-93690c80-e184-11eb-8023-7dfb78eaefa2.png">

