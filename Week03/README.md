#### 一、errgroup

1.执行完所有任务再返回，直接使用 errgroup.Group 结构

2.如果期望遇到错误立即停止所有任务，则使用 errgroup.WithContext

