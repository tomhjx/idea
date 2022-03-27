# Idea
模拟一些场景，给出评估及意见

# Logic

* 输入
    * 描述目标行为及期望结果
        * e.g. 行为（调用xxx接口，用户在页面下单，用户在app下单）；期望结果（QPS：xxx, 响应时间：xxx MS）
    * 描述资源供给
        * e.g. 带宽、cpu、内存、硬盘
* 评估
    * 解析配置
        * 引用组件库
    * 构建架构
    * 实例化组件对象，编排组件
    * 构建执行链路
    * 应用调度策略
    * 依赖用例库进行演练
        * 常规场景下推演
        * 异常场景下推演

* 输出
    * 建议选型
    * 资源需求

# Usage

```bash

idea -e=/path/to/expect.yaml -r=/path/to/resource.yaml

// output e.g. {"cpu_total":12, "memory_total":"110GB"}
```