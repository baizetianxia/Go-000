## 学习笔记 作业
问：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

课程内容：
调用标准库或第三方库(github库)或公司基础库，产生的error,考虑使用errors.Wrap或Wrapf保存堆栈信息，向上传递至程序顶部，再使用%v吧堆栈详情记录下来。

答：dao层对数据库进行操作，遇到sql.ErrNoRows的error，应该将该error向上抛。service层将调用dao层的函数时抛上来的error，通过Wrap包装并增加新的信息，然后层层上抛，向上传递至程序顶部，再使用%v吧堆栈详情记录下来。
