# TypeScript 基本语法

[TOC]

## After all

TypeScript（简称 TS）是由微软开发并维护的一种开源编程语言，它是 JavaScript 的一个超集（superset），在 JavaScript 的基础上增加了**静态类型系统**和其他现代语言特性。TypeScript 最终会被编译为纯 JavaScript 代码，可以在任何支持 JavaScript 的环境中运行。

TypeScript文件定义为`xxx.ts`，可使用`tsx note.ts`命令行运行，需要先安装tsx：`npm install -g tsx`

## 基础类型

TypeScript支持与JavaScript几乎相同的数据类型，同时还支持枚举类型。

### 布尔值 boolean

```typescript
let isOK: boolean = false;
```

### 数字 number

```typescript
// 和 JS 一样，TS 不区分整数和浮点数，所有数字都是浮点数。除十进制外还支持二进制八进制十六进制字面量。
let decNumber: number = 100;
let hexNumber: number = 0xf00d;
let binaryNumber: number = 0b1010;
let octalNumber: number = 0o744;
```

### 字符串 string

```typescript
// TS 中可以使用双引号（""）或者单引号（''）表示字符串。
let str: string = "Hello, World!";
str = "Hi, Benjamin!";
// TS 原生支持模板字符串，模板字符串需要用反引号（``）包围，在其内部使用（${}）引用变量，也可以嵌入表达式
console.log(`${str} Good morning!`);
// 字符串可以使用（+）进行拼接
console.log(str + " Good morning!");
```

### 数组

```typescript
// 表示相同类型的若干个数字，初始化时元素数量就固定了，可以更改但不能增减。
let list: number[] = [1, 2, 3]
let arr: Array<number> = [1, 2, 3]
```

### Map

```typescript
// Map是JS的ES6标准引入的对象，并不是TS提供的
const m = new Map([
  ["a", 1],
  ["b", 2],
  ["c", 3]
]); // Map<string, number>
m.set("d", 4)
console.log(m.get("a")); // 1
```



### 元组

```typescript
// 元组类型表示一个已知元素数量和元素类型的数组，各元素的类型不必相同，值也可以更改但必须符合类型。
let x: [string, number] = ["abc", 123]
x = [123, "abc"] // error
```

### 枚举 enum

```typescript
// enum 类型是 TS 对 JS 的补充，可以定义一系列名字以及其对应的元素编号
// 不主动声明的情况下，从 0 开始编号
enum paymentType {Cash, CreditCard, Mobile} // 等同于{Cash = 0, CreditCard = 1, Mobile = 2}
console.log(paymentType.Cash) // output 1
// 也可以自定义起始编号或者自定义每一个编号
enum paymentType {Cash = 1, CreditCard, Mobile}
enum paymentType {Cash = 2, CreditCard = 4, Mobile = 6, Others = 999}
// 我们可以通过枚举的编号获取其名称
console.log(paymentType[2]) // output "Cash"
```

### any

```typescript
// any 表示在编程阶段暂时不知道这个变量的类型，可能来自于动态的内容，比如用户输入或者三方库返回。在编译阶段我们不希望类型检查器对这些变量进行类型检查，可以使用 any 定义其类型为‘任何类型’。
// 与 JS 的 Object 相比，any 类型的对象可以调用其真的拥有的方法，而 Object 只允许赋值，不允许调用方法，即便它真的拥有这个方法。
let notSure: any;
notSure = 1;
notSure = "str";
console.log(notSure.length); // output "3"
```

### void

```typescript
// 某种程度上来说，void类型与any类型相反，void表示没有任何类型。当一个返回值没有返回值时，可以定义为void
function aLog(): void {
    console.log("a log");
}
// 声明一个 void 类型的变量其实没啥大用，因为它只能赋值 undefined 或者 null。
let u: void = undefined;
```

### null 和 undefined

```typescript
// null 和 undefined 表示没有值，他们分别有自己的类型 null 和 undefined。
// 默认情况下 null 和 undefined 是所有类型的子类型，可以将 null 和 undefined 赋值给任何类型的变量。
// 当使用了 --strictNullChecks 标记是，null 和 undefined 只能赋值给 void 和他们各自类型的变量。鼓励使用 --strictNullChecks，可以避免很多常见问题。
// 当使用了 --strictNullChecks 时又想给某个元素赋值 null 或者 undefined 时，可以使用联合类型 string | null | undefined
```

### never

```typescript
// never 类型表示永远不会发生的值，比 void 更严格。
// 它是所有类型的子类型，但没有任何类型是 never 的子类型（除了它自己）。
// 常用于：
// 抛出异常的函数，抛出错误，不会返回任何值
function throwError(message: string): never {
  throw new Error(message);
}
// 无限循环的函数
function infiniteLoop(): never {
  while (true) {
    // do something
  }
}
// 类型收窄（Type Narrowing）中的“穷尽性检查”
type Color = 'red' | 'green' | 'blue';

function getColorName(c: Color): string {
  switch (c) {
    case 'red': return '红色';
    case 'green': return '绿色';
    case 'blue': return '蓝色';
    default:
      // 如果漏掉某个 case，这里 c 的类型会是 never
      const _exhaustiveCheck: never = c;
      throw new Error(`未处理的颜色: ${_exhaustiveCheck}`);
  }
}
```

### object

```typescript
// object表示非原始类型，也就是除 number，string，boolean，symbol，null，undefined 之外的类型
```

### 类型断言

```typescript
// 对于 any 类型的变量，有时候我们清楚明白它的类型时，可以显示的将它当作这类型去使用
// 有两种方式，分别是尖括号和 as
let aValue: any = "a value";
console.log((<string>aValue).length);
console.log((aValue as string).length);
```

## 变量声明

### var 声明

在JavaScript中我们总是通过var关键字来定义变量。

```typescript
var aValue = 10;
```

使用 `var` 声明的变量：

-   在**整个函数内**都可见；

-   即使在 `if`、`for`、`while` 等**代码块内部声明**，也会**提升（hoisted）到函数顶部**，但**只有声明被提升，赋值不会**。

    -   ```typescript
        console.log(a); // undefined（不是报错！）
        var a = 5;
        console.log(a); // 5
        
        // 等价于
        var a;          // 声明被提升
        console.log(a); // undefined
        a = 5;          // 赋值留在原地
        console.log(a); // 5
        ```

-   如果不在任何函数内，则成为**全局变量**。

-   可以重复声明而不报错。

TypeScript建议**不要使用 `var`**，所以可以直接忽略var直接学习let和const。

### let 声明

let声明和var声明的写法一致。

```typescript
let aValue = 10;
```

let的特性：

-   块级作用域：let声明的变量**只在代码块`{}`内部有效**。（包括for循环）
-   **不会被提升到块顶部**，在声明之前访问会报错。
-   同一作用域内**不允许重复声明**。

### const 声明

const声明用于声明常量，声明时必须同时进行初始化，并且初始化后不可以再次进行赋值。

```typescript
const name; // error
const name = "Benjamin" // ✅
name = "Ben" // error
```

const声明出来的常量拥有与let声明相同的作用域。

当const声明出来一个引用类型时，该引用不可改，但是引用变量内部的成员是可以改的。幸运的是TypeScript允许将对象的成员设置成只读的。

```typescript
const people = {
    name: "Ben",
    age: 26,
}
people.name = "Benjamin" // ✅
```

### let vs const

既然let和const的作用域相同，那么如何进行选择呢？

根据最小特权原则，所有变量都尽可能优先使用const，当已知会对其进行修改时，才使用let进行声明。

### 解构

**解构（Destructuring）** 是一种从数组或对象中**提取值并赋给变量**的简洁语法。它让代码更清晰、更易读，是现代 JS/TS 开发的核心特性之一。

**数组解构：**

```typescript
const colors = ['red', 'green', 'blue'];

// 按位置提取
const [first, second, third] = colors;
console.log(first);  // 'red'
console.log(second); // 'green'

// 跳过元素
const [,, third] = colors; // 跳过前两个
console.log(third); // 'blue'

// 默认值
const [a, b = 'default'] = ['hello'];
console.log(b); // 'default'
const [a, b = 'default'] = ['hello', 'world'];
console.log(b); // 'world'

// 剩余元素
const [head, ...tail] = [1, 2, 3, 4];
console.log(head); // 1
console.log(tail); // [2, 3, 4]

// 交换元素
let x = 1, y = 2;
[x, y] = [y, x];
console.log(x, y); // 2, 1
```

**对象解构：**

```typescript
const user = { name: 'Alice', age: 30 };

// 按属性名提取，变量名需与属性名一致
const { name, age } = user; // name='Alice', age=30

// 重命名
const { name: fullName, age: years } = user;
console.log(fullName); // 'Alice'

// 默认值
const { name, role = 'guest' } = { name: 'Bob' };
console.log(role); // 'guest'

// 嵌套解构
const person = {
  name: 'Charlie',
  address: {
    city: 'Beijing',
    country: 'China'
  }
};
const { name, address: { city } } = person;
console.log(city); // 'Beijing'
```

**函数声明解构：**

我的理解是当使用param这种符合结构体进行传参时，函数参数解构可以对其直接解析成若干变量，而不是在函数内部使用`.`进行访问，而且可以显示的声明可能为空时设置的默认值。

```typescript
function greet({ name, greeting = 'Hello' }: { name: string; greeting?: string }) {
  console.log(`${greeting}, ${name}!`);
}

greet({ name: 'David' }); // "Hello, David!"
```

### 展开

“**展开**”（Spread）通常指使用 **展开运算符（Spread Operator）** `...`，它可以将**可迭代对象（如数组、对象等）“打散”为独立元素或属性**。这是 ES6+ 和现代 TS 开发中的核心语法之一。

**展开数组：**将一个数组展开为多个元素。

```typescript
const arr1 = [1, 2, 3];
const arr2 = [4, 5, 6];

// 浅拷贝复制数组
const copy = [...arr1]; // [1, 2, 3]

// 合并数组
const merge = [...arr1, ...arr2] // [1, 2, 3, 4, 5, 6]

// 在函数调用中展开参数，不需要再去一个一个点了
Math.max(...arr1); // 等价于 Math.max(arr1[1], arr1[2], arr1[3])
```

**展开对象：**将一个对象的可枚举属性“复制”到新对象中。

```typescript
const config1 = { theme: 'light', lang: 'en' };
const config2 = { lang: 'zh', debug: true };


// 复制对象（浅拷贝）
const copy_config1 = { ...config1 };

// 合并对象，若后面对象与前面对象有同名属性，后面的优先级高
const merge_config = { ...config1, ...config2 }; // { theme: 'light', lang: 'zh', debug: true }

// 新增/覆盖属性
const new_config = { ...merge_config, debug: false, dry_run: true }; // { theme: 'light', lang: 'zh', debug: false, dry_run: true }

```

**在函数参数中收集参数：**

```typescript
// 聚合参数
function sum(...numbers: number[]) {
  return numbers.reduce((a, b) => a + b, 0);
}
sum(1, 2, 3, 4); // numbers = [1, 2, 3, 4]

// 收集剩余参数
function sumFirstAndSecond(first: number, second: number, ...others: number[]) {
  return first + second;
}
sumFirstAndSecond(1, 2, 3, 4); // 3
```

## 运算符

TypeScript的大部分运算符都和主流强类型语言相同，这里简单提一下，只介绍其独特的部分。

```typescript
// typeof 运算符
// 是一元运算符，返回操作数的数据类型。
var num = 12 
console.log(typeof num);   //输出结果: number

// instanceof 运算符
// 用于检测一个对象是否是某个构造函数（或类）的实例
class Animal {}
class Dog extends Animal {}
const dog = new Dog();
dog instanceof Dog;    // true
dog instanceof Animal; // true（继承关系）
dog instanceof Object; // true（所有对象最终继承自 Object）
```

## 流程控制语句

### 条件语句

```typescript
// if 语句
if(expression 1) {
    // do something
} else if(expression 2) {
    // do something
} else {
    // do something
}

// switch…case 语句
// 与Java中switch语句语法一致
switch(expression){
    case constant-expression:
       // do something
       break;
    case constant-expression:
       // do something
       break;
    default:
       // do something
}
```

### 循环语句

**for循环**

```typescript
// for循环
// 适用于需要索引控制或高性能场景（如大型数组）
for(let i = 0; i < 10; i++) {
    console.log(i);
}

// for in循环可以遍历可迭代对象的可枚举属性名（key）
// 不建议使用于数组
let chars: string[] = ["a", "b", "c", "d", "e"];
for(let i in chars) {
    console.log(i); // 0 1 2 3 4
}
// for of循环可以遍历可迭代对象的子元素 ✅建议使用
let chars: string[] = ["a", "b", "c", "d", "e"];
for(let c of chars) {
    console.log(c); // a b c d e
}
```

**数组高阶方法 ✅ 函数式风格，强烈推荐**

```typescript
// forEach
[1, 2, 3].forEach(num => {
  // num: number
});

// map 对迭代对象的每一个元素调用函数
const doubled = [1, 2, 3].map(x => x * 2); // [ 2, 4, 6 ]

// filter 收集调用函数返回真的元素
const filtered = [1, 2, 3].filter(x => x % 2 === 0); // [2]

// reduce 用于将数组“归约”为一个单一的值
// 会对迭代对象的每一个元素执行第一个参数上的函数
// acc是上一次执行的返回值，x是当前迭代元素
// 第二个参数是acc的初始值
const sum = [1, 2, 3].reduce((acc, x) => acc + x, 0); // 6

```

### 遍历对象的键值

不要用 `for...in`，改用：

✅ `Object.keys()` / `Object.values()` / `Object.entries()`

```typescript
const user = { name: "Alice", age: 30 } as const;

// keys: ("name" | "age")[]
Object.keys(user).forEach(key => {
  // key: string（需断言）
});

// 更安全：使用 Object.entries + as const
for (const [key, value] of Object.entries(user)) {
  // key: "name" | "age", value: "Alice" | 30
}
```



## 类型定义

### class

在 TypeScript 中，`class` 是用于定义类的关键字，它允许你使用面向对象的方式组织代码。TypeScript 的类是对 ES6（ECMAScript 2015）类的增强，增加了类型注解、访问修饰符（如 `public`、`private`、`protected`）、抽象类等功能。

基本等同于java中的class。

```typescript
class Person {
  // 属性（字段）
  private name: string;
  protected age: number;

  // 构造函数
  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }

  // 方法
  greet(): string {
    return `Hello, my name is ${this.name} and I am ${this.age} years old.`;
  }

  // Getter
  getName(): string {
    return this.name;
  }

  // Setter
  setName(newName: string): void {
    this.name = newName;
  }
}

// 使用类
const person = new Person("Alice", 30);
console.log(person.greet()); // Hello, my name is Alice and I am 30 years old.
```

**访问修饰符说明**：

-   `public`（默认）：可以在**任何地方**访问。
-   `private`：只能在**类内部**访问。
-   `protected`：可以在**类内部及其子类**中访问。

**继承**:

```typescript
class Employee extends Person {
  private department: string;

  constructor(name: string, age: number, department: string) {
    super(name, age); // ⚠️在子类构造函数中调用this前，必须先调用父类构造函数
    this.department = department;
  }

  getDetails(): string {
    // 注意：不能访问父类的 private 成员 name，但可以访问 protected 成员 age
    return `${this.getName()} works in ${this.department}.`;
  }
}

const emp = new Employee("Bob", 25, "Engineering");
console.log(emp.getDetails()); // Bob works in Engineering.
```

**抽象类**：

```typescript
abstract class Animal {
  abstract makeSound(): void;

  move(): void {
    console.log("Moving...");
  }
}

class Dog extends Animal {
  makeSound(): void {
    console.log("Woof! Woof!");
  }
}

// const animal = new Animal(); // 错误：不能实例化抽象类
const dog = new Dog();
dog.makeSound(); // Woof! Woof!
dog.move();      // Moving...
```

**其他特性**：

-   **静态成员**：使用 `static` 关键字定义属于类本身而不是实例的属性或方法。
-   **只读属性**：使用 `readonly` 关键字定义初始化后不可更改的属性。
-   **参数属性**：在构造函数中直接声明并初始化属性，例如：

```typescript
class Point {
  constructor(public x: number, public y: number) {}
}
// 等价于：
// class Point {
//   public x: number;
//   public y: number;
//   constructor(x: number, y: number) {
//     this.x = x;
//     this.y = y;
//   }
// }
```

### interface

在 TypeScript 中，`interface`（接口）是一种非常核心的类型定义机制，用于描述对象的结构（形状），包括属性、方法等。它**只在编译时存在**，不会出现在最终的 JavaScript 代码中（即“零成本抽象”）。

与class的区别是只**描述“形状”（契约）**，不提供实现。class需要提供实现，否则需要定义为抽象类。

>   ✅ 简单说：
>
>   -   `interface` 告诉你“一个东西**长什么样**”，是**“类型契约”**。
>   -   `class` 告诉你“一个东西**是什么，并且能做什么**”，是**“可执行的实体”**。
>
>   在设计时：先想清楚你是要**约束结构**（用 interface），还是要**创建对象**（用 class）。

**基本语法**

```typescript
interface User {
  name: string;
  age: number;
  isActive?: boolean; // 可选属性（加 ?）
}
const alice: User = {
  name: "Alice",
  age: 30
  // isActive 可省略
};
```

**核心特性**

```typescript
// 可选属性
interface Config {
  theme: string;
  lang?: string; // 可有可无
}

// 只读属性
interface Point {
  readonly x: number;
  readonly y: number;
}
const p: Point = { x: 10, y: 20 };
p.x = 5; // ❌ Error! 只读属性不可以赋值

// 函数类型
interface GreetFunction {
    (name: string): string;
}
const greet: GreetFunction = function (name: string) {
    return `Hello, ${name}`;
}
console.log(greet("Benjamin"));

// 方法声明
interface Person {
  name: string;
  greet(): void; // 方法
}
class Student implements Person {
  name = "Bob";
  greet() {
    console.log("Hi!");
  }
}
```

**高级能力**

```typescript
// 继承：接口可以继承一个或多个其他接口。
interface Flyable { fly(): void; }
interface Swimmable { swim(): void; }
interface Duck extends Flyable, Swimmable {
  quack(): void;
}

// 合并：同名接口会自动合并（这是 `interface` 相比 `type` 的独特优势）。
interface User {
  name: string;
}
interface User {
  age: number;
}
// 合并后等价于：
// interface User {
//   name: string;
//   age: number;
// }

// 实现类：类可以用 `implements` 关键字承诺实现某个接口。
interface ClockInterface {
  currentTime: Date;
  setTime(d: Date): void;
}
class Clock implements ClockInterface {
  currentTime: Date = new Date();
  setTime(d: Date) {
    this.currentTime = d;
  }
}
```

### type

在 TypeScript 中，`type`关键字用于**创建类型别名（Type Alias）**。它允许你为任意类型（包括原始类型、联合类型、元组、对象类型、函数类型，甚至复杂的泛型类型）定义一个可重用的名称。

与interface的区别在于：type只是个类型别名，强调类型；而interfase更强调结构，希望代码面向对象

>   -   **对象结构用 `interface`，类型组合用 `type`**
>   -   **要合并、要继承 → `interface`**
>   -   **要联合、要元组、要泛型魔法 → `type`**
>   -   二者都可的情况下优先使用interface。 “**Prefer `interface` over `type` for object types unless you need specific features of `type`.**”

**基本语法**

```typescript
type TypeName = TypeDefinition;
```

其中：

-   `TypeName` 是你自定义的类型名称（通常首字母大写）。
-   `TypeDefinition` 可以是任何有效的 TypeScript **类型**表达式。

```typescript
// 基本类型别名
type ID = string | number;
let userId: ID = "abc123";
userId = 42; // ✅ 合法

// 对象类型
type User = {
  name: string;
  age: number;
};
const user: User = { name: "Alice", age: 30 };

// 联合类型
type Status = "idle" | "loading" | "success" | "error";
let status: Status = "loading";

// 元组
type Point = [number, number];
const p: Point = [10, 20];

// 函数
type AddFn = (a: number, b: number) => number;
const add: AddFn = (x, y) => x + y;

// 泛型类型别名
type Response<T> = {
  data: T;
  status: number;
};
const userResponse: Response<User> = {
  data: { name: "Bob", age: 25 },
  status: 200
};
```

### `interface` vs `class` vs `type`

**`interface`----- “形状”的契约**

-   **描述结构**，不包含实现。
-   **支持合并**（同名自动合并）。
-   **可被 class 实现**（`implements`）。
-   **不能表示原始类型、联合类型等**。

>   📌 适用：API 设计、组件 Props、配置对象等**公共契约**。

**`class` ----- 可实例化的实体**

-   **同时是值（构造函数）和类型（实例类型）**。
-   包含**具体实现**（方法体、初始化逻辑）。
-   可以 `implements interface` 或 `extends class`。

>   📌 适用：需要**封装状态 + 行为**的对象（如模型、服务类）。

**`type` ----- 灵活的类型别名**

-   **可以表示任何类型**（包括 `interface` 能表示的）。
-   **支持联合、元组、映射类型等高级特性**。
-   **不能被 `implements` 或 `extends`**（因为它可能不是对象类型）。
-   **不支持合并**。

>   📌 适用：工具类型、状态枚举、复杂类型组合。

## 函数

在 **TypeScript（TS）** 中，**函数（Function）** 不仅包含运行时逻辑，还带有**完整的类型信息**，包括参数类型、返回值类型、可选参数、重载等。TS 的函数系统是 JavaScript 函数的超集，提供了强大的类型安全和开发体验。

### 函数的基本类型

**普通函数**

```typescript
function greet(name: string): string {
  return `Hello, ${name}!`;
}
```

**箭头函数**

```typescript
const add = (a: number, b: number): number => a + b;
```

### 参数相关特性

**可选参数**

使用`?`标注的参数被认为是可选参数，在传参时可以不传，不传会被认为是说`undefined`。可选参数必须放在必传参数后面

```typescript
function greet(name: string, title?: string): string {
  return title ? `${title} ${name}` : name;
}
greet("Alice");        // OK
greet("Bob", "Dr.");   // OK
```

**默认参数**

有默认值的参数**自动成为可选参数**。默认参数可以不写类型，TypeScript会自动推断类型。

```typescript
function greet(name: string, greeting: string = "Hello"): string {
  return `${greeting}, ${name}!`;
}
```

**剩余参数**

`...`展开符可以收集剩余参数，需放在所有明确定义的参数（包括必传，可选参数）的后面。

剩余参数可以匹配不到，可以匹配一个，也可以匹配多个，但不管怎样，都会用数组的形式给出。

```typescript
function buildName(firstName: string, ...restOfName: string[]): string {
  return firstName + " " + restOfName.join(" ")
}

buildName("Franklin", "Delano", "Roosevelt"); // Franklin Delano Roosevelt
```

### 函数类型

使用`function`关键字时需要提供函数的实现（就算不立刻给出，也需要在下文中定义），很多情况下我们需要仅定义函数的名称参数返回值（将三者合并统称为类型），仅定义函数类型的方法有很多：

**使用接口 interface**

```typescript
interface AddFunc {
  (a: number, b: number): number;
}
const add: AddFunc = (x, y) => x + y;
```

**使用类型别名 type （更常见）**

```typescript
type AddFunc = (a: number, b: number) => number;
const add: AddFunc = (x, y) => x + y;
```

**内联注解 （常用于回调）**

```typescript
function process(callback: (result: string) => void) {
  callback("done");
}
```

### 函数重载

当一个函数需要根据**不同参数类型返回不同结果**时，使用重载：

```typescript
// 重载签名（多个）
function combine(a: string, b: string): string;
function combine(a: number, b: number): number;

// 实现签名（一个，必须兼容所有重载）
function combine(a: string | number, b: string | number): string | number {
  if (typeof a === 'string' && typeof b === 'string') {
    return a + b;
  }
  if (typeof a === 'number' && typeof b === 'number') {
    return a + b;
  }
  throw new Error('Invalid arguments');
}

// 使用
combine("Hello", "World"); // 返回 string
combine(10, 20);           // 返回 number


// 一个典型错误，下面这个方法在参数a和b都是number类型时，返回了string，是不符合重载的两种签名的，但是TypeScript并不会检查函数实现中的具体逻辑，所以代码检查以及编译阶段是不会报错的，开发者需要手动保持一致性。
// 关键原则：函数重载的实现必须满足每一个重载签名的契约。
function combine(a: string | number, b: string | number): string | number {
  if (typeof a === 'string' && typeof b === 'string') {
    return a + b;
  }
  if (typeof a === 'number' && typeof b === 'number') {
    return "aaaaa";
  }
  throw new Error('Invalid arguments');
}
```

### this 类型控制

TypeScript 允许你在**函数签名的第一个参数位置**写一个伪参数 `this`，用于声明该函数被调用时 `this` 应该是什么类型。

>   个人理解：
>
>   this是一个方法内部的隐式对象，表示调用此方法的对象。方法内部可以通过this获取调用者的信息
>
>   在方法签名的第一个参数定义this的类型可以定义什么对象才可以调用这个方法。

```typescript
interface User {
  name: string;
  greet(this: User): void;
}

function greet(this: User) {
  console.log(`Hello, I'm ${this.name}`);
}

const user: User = { name: "Alice", greet };

// 正确调用（this 是 User）
user.greet(); // ✅ OK

// 错误调用（this 上下文丢失）
const fn = user.greet;
fn(); // ❌ 编译错误！
// Error: The 'this' context of type 'void' is not assignable to method's 'this' of type 'User'.
```

### 泛型函数

让函数支持多种类型，同时保持类型安全：

```typescript
function identity<T>(arg: T): T {
  return arg;
}

const str = identity<string>("hello"); // T = string
const num = identity(42);              // T = number（自动推断）
```

带约束的泛型：

```typescript
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}

const person = { name: "Alice", age: 30 };
getProperty(person, "name"); // 类型：string
getProperty(person, "email"); // ❌ Error! "email" 不是 person 的属性
```

### 常见函数类型工具

TS 内置了一些实用的函数相关工具类型：

| 工具类型                   | 作用                 |
| -------------------------- | -------------------- |
| `Parameters<T>`            | 提取函数参数类型元组 |
| `ReturnType<T>`            | 提取函数返回值类型   |
| `ConstructorParameters<T>` | 提取构造函数参数     |

```typescript
function fetchData(url: string, timeout: number): string { 
    return url + timeout;
 }

type FetchParams = Parameters<typeof fetchData>; // [string, number]
type FetchResult = ReturnType<typeof fetchData>;  // string
```

## 泛型

TypeScript 中的 **泛型（Generics）** 是其类型系统最强大、最核心的特性之一。它允许你编写**可复用、类型安全且灵活**的组件（函数、类、接口等），而无需提前指定具体类型——而是让使用者在调用时决定类型。

在不使用泛型的情况下，我们在编程过程中想表示“某种类型”只能使用`any`关键字，但是多次使用`any`可能会丢失前后的类型一致性。使用泛型我们则可以通过一个模拟的类型`T`代表“这个类型”。

>   💡 **泛型的核心思想**：**“类型作为参数”** —— 在使用时传入类型。

### 基本语法

**函数泛型**

```typescript
function firstElement<T>(arr: T[]): T | undefined {
  return arr[0];
}

firstElement([1, 2, 3]);     // 返回 number | undefined
firstElement(["a", "b"]);    // 返回 string | undefined
```

**泛型接口**

```typescript
interface Box<T> {
  value: T;
}

const stringBox: Box<string> = { value: "hello" };
const numberBox: Box<number> = { value: 42 };
```

**泛型类**

```typescript
class Queue<T> {
  private items: T[] = [];

  enqueue(item: T) {
    this.items.push(item);
  }

  dequeue(): T | undefined {
    return this.items.shift();
  }
}

const queue = new Queue<string>();
queue.enqueue("task1"); // OK
queue.enqueue(123);     // ❌ Error!
```

**泛型类型别名**

```typescript
type Pair<T> = [T, T];

const p: Pair<number> = [1, 2];
```

### 泛型约束

**使用 `extends` 约束**

```typescript
interface HasLength {
  length: number;
}

function logLength<T extends HasLength>(arg: T): T {
  console.log(arg.length); // 安全！因为 T 满足 HasLength 接口，必须有 length
  return arg;
}

logLength("hello");   // OK（string 有 length）
logLength([1, 2, 3]); // OK（array 有 length）
logLength(42);        // ❌ Error! number 没有 length
```

**常见约束：`keyof`**

```typescript
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key]; // 类型安全！
}

const person = { name: "Alice", age: 30 };
getProperty(person, "name"); // 返回 string
getProperty(person, "email"); // ❌ 编译错误！
```

### 泛型默认值

可以为泛型参数提供默认类型（类似函数默认参数）：

```typescript
interface Response<T = string> {
  data: T;
  status: number;
}

// 不指定 T 时，默认为 string
const res1: Response = { data: "success", status: 200 };

// 显式指定
const res2: Response<number> = { data: 42, status: 200 };
```

### 高级用法

**多个泛型参数**

```typescript
function zip<A, B>(a: A[], b: B[]): [A, B][] {
  return a.map((item, i) => [item, b[i]]);
}

zip([1, 2], ["a", "b"]); // ([number, string])[]
```

**泛型 + 条件类型（Conditional Types）**

```typescript
type IsString<T> = T extends string ? true : false;

type A = IsString<"hello">; // true
type B = IsString<42>;      // false
```

**工具类型**

```typescript
// Partial<T>：将 T 的所有属性变为可选
type PartialUser = Partial<{ name: string; age: number }>;
// 等价于 { name?: string; age?: number }

// Pick<T, K>：从 T 中选出 K 中的属性
type NameOnly = Pick<{ name: string; age: number }, "name">;
// { name: string }
```

## 导入导出

在 TypeScript（TS）中，`import` / `export` 是 **ES6 模块系统（ECMAScript Modules, ESM）** 的核心语法，也是 **现代 JavaScript 和 TS 推荐的模块化方式**。它用于**组织代码、封装逻辑、管理依赖**，并提供**完整的类型安全支持**。

>   只要一个 TS 文件包含 `import` 或 `export` 语句，它就被视为一个 **模块（Module）**，拥有自己的作用域（不会污染全局）。

**普通导入导出**

```typescript
// Math.ts
// 直接导出
export const PI = 3.14159;
export function add(a: number, b: number): number {
  return a + b;
}
// 也可以先定义，再统一导出
const multiply = (a: number, b: number) => a * b;
export { multiply };


// Other.ts
import { PI, add, multiply } from './Math';
```

**默认导入导出**

```typescript
// User.ts
class User {
  constructor(public name: string) {}
}

export default User; // 一个文件只能有一个 default export，默认导出不用加大括号

// Other.ts
import User from './User'; // 直接用一个对象承载，且名字可自定义
const user = new User("Alice");
```

