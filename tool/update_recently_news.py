# -*- coding: utf-8 -*-

import requests
def main():
    response = requests.get('https://www.city.hamamatsu.shizuoka.jp/koho2/emergency/korona.html')
    response.encoding = response.apparent_encoding
    print(response.text)

if __name__ == '__main__':
    main()