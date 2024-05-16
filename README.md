# medium-to-md

medium 目前导出的策略是将回复跟文章都放在posts文件夹中
目前这个程序只处理了posts下 大小超过 5000B 的 文件
--- 
你可以在template 自定义 文件输出的格式
遗留问题
1. 在原始文件中，<footer> 标签里面有createTime 需要取出来
2. 代码块其实有标注代码类型，也可以弄出来。
3. 需要支持自定义图床上传

# Third-party libraries
1. github.com/JohannesKaufmann/html-to-markdown  主要的转换逻辑
2. github.com/hxzhouh/html-to-markdown 从 github.com/JohannesKaufmann/html-to-markdown fork 后续会做一些针对medium 的定制

