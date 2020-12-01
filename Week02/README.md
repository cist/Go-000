##学习笔记
日志记录与错误无关且对调试没有帮助的信息应被视为噪音，应予以质疑。记录的原因是因为某些东西失败了，而日志包含了答案。

错误要被日志记录。
应用程序处理错误，保证100%完整性。
之后不再报告当前错误。

##作业：
1.我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？ 

答：不应该返回sql.ErrNoRows。如果上层需要使用当sql.ErrNoRows，导致当前业务层依赖外部库，不利于后面扩展修改。同时上层也不应该看懂sql的对应错误信息。
建议dao层应当根据当前的业务逻辑和与上层的接口定义，封装一个符合业务逻辑并且定义好的err，再wrap出去。

## 课后查看连接
https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right
https://dave.cheney.net/2015/01/26/errors-and-exceptions-redux
https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux
https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html
https://morsmachine.dk/error-handling
https://blog.golang.org/error-handling-and-go
https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html
https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html
https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html
https://blog.golang.org/errors-are-values
https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package
https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html
https://crawshaw.io/blog/xerrors
https://blog.golang.org/go1.13-errors
https://medium.com/gett-engineering/error-handling-in-go-53b8a7112d04
https://medium.com/gett-engineering/error-handling-in-go-1-13-5ee6d1e0a55c




https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
