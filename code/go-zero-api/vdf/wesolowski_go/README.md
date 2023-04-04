x### VDF (verified delay function) implementation in Golang


|    名称    | 已实现  | 来源                                   | 描述                       |
| :-------: | :----: | :------------------------------------ | -------------------------- |
|    vdf    |   √    | `https://github.com/harmony-one/vdf`  | harmony链vdf包              |

本部分针对密码学库中的验证延迟函数模块进行详细说明。验证延迟函数采用Wesolowski的VDF，针对一个输入，可以连续计算得到输出和一个验证输出的证明，计算结果不可以并行加速，利用证明可以快速验证计算的正确性。

In this implementation, the VDF function takes 32 bytes as seed and an integer as difficulty.   

Please note that only 2048 integer size for class group variables are supported now.  