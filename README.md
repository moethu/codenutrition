# Code Nutrition Facts

[![nutrition facts](http://82.165.112.45:5000/badge/O%2B%2B_S%2B%2B_I%2B%2B_C_E_M_V%2B_PS%2B%2B_D%2B)](http://82.165.112.45:5000/facts/O%2B%2B_S%2B%2B_I%2B%2B_C_E_M_V%2B_PS%2B%2B_D%2B)

**Imagine your code and all dependencies would carry nutrition labels.**

Code Nutrition Facts is a service for labelling code so everybody can understand what they are getting into when building upon or using your software. The survey is based on a set of questions presented by Felix von Leitner: "Nützlich-Unbedenklich Spektrum" at 36C3 (Minor adjustmenst have been made to the original text.) [learn more here](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=1&cad=rja&uact=8&ved=2ahUKEwjE-v7ropLnAhXUwMQBHd68B9UQwqsBMAB6BAgKEAQ&url=https%3A%2F%2Fmedia.ccc.de%2Fv%2F36c3-10608-das_nutzlich-unbedenklich_spektrum&usg=AOvVaw1_05ix3-K_lRn_T9LbJRZi).

The survey itself is still under development you are welcome to discuss, contribute and improve it. Currently it consists of 9 fundamental questions regarding the state of your (legacy) code. After answering them by multiple choice the service generates an embeddable badge representing the state of your project. I highly reccommend completing the survey even if you don't want to use the badge - the questions asked can be very helpful.

[Survey](https://github.com/moethu/codenutrition/blob/master/static/spectrum.json)

Again, the intention of this repo and the survey is to get the conversation started and potentially develop a solid metric representing the state of a project.

[Start your survey now](http://82.165.112.45:5000)

## Requirements for building the service

- go 1.11.5+

### requirements (via go modules)

- github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
- github.com/gin-gonic/gin v1.5.0
- github.com/llgcode/draw2d v0.0.0-20200110163050-b96d8208fcfc
- github.com/swaggo/gin-swagger v1.2.0
- github.com/swaggo/swag v1.6.5
