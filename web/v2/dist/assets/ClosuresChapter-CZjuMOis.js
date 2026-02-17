import{c as e,d as t}from"./index-BWDKUK_q.js";var n=t();function r(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsxs)(`p`,{children:[(0,n.jsx)(`b`,{children:`Замыкание (Closure)`}),` — это комбинация функции и лексического окружения, в котором эта функция была определена.`]}),(0,n.jsx)(`p`,{children:`Простыми словами: замыкание позволяет функции запоминать и иметь доступ к переменным из своей внешней области видимости даже после того, как внешняя функция завершила своё выполнение.`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Как это работает?`}),(0,n.jsx)(`p`,{children:`В JavaScript функции являются объектами «первого класса». Это означает, что их можно передавать как аргументы, возвращать из других функций и присваивать переменным.`}),(0,n.jsx)(e,{code:`function createCounter() {
  let count = 0; // Переменная во внешней области видимости

  return function() {
    count++; // Функция "запоминает" переменную count
    return count;
  };
}

const counter = createCounter();

console.log(counter()); // 1
console.log(counter()); // 2
console.log(counter()); // 3`}),(0,n.jsxs)(`p`,{children:[`В этом примере `,(0,n.jsx)(`code`,{children:`counter`}),` — это замыкание. Оно состоит из внутренней функции и переменной `,(0,n.jsx)(`code`,{children:`count`}),` из области видимости `,(0,n.jsx)(`code`,{children:`createCounter`}),`.`]}),(0,n.jsx)(`p`,{children:`Когда ты делаешь:`}),(0,n.jsx)(e,{code:`const counter = createCounter();`}),(0,n.jsx)(`h3`,{children:`Шаги`}),(0,n.jsxs)(`ol`,{children:[(0,n.jsxs)(`li`,{children:[`Вызывается `,(0,n.jsx)(`code`,{children:`createCounter`})]}),(0,n.jsxs)(`li`,{children:[`Создаётся переменная `,(0,n.jsx)(`code`,{children:`count = 0`})]}),(0,n.jsxs)(`li`,{children:[`Возвращается `,(0,n.jsx)(`b`,{children:`внутренняя функция`})]}),(0,n.jsxs)(`li`,{children:[`ВАЖНО: эта функция `,(0,n.jsx)(`b`,{children:`сохраняет ссылку на переменную`}),` `,(0,n.jsx)(`code`,{children:`count`})]})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Зачем нужны замыкания?`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`b`,{children:`Инкапсуляция и приватные переменные:`}),` скрыть данные от прямого доступа извне.`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`b`,{children:`Создание фабричных функций:`}),` генерация функций с предустановленным поведением.`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`b`,{children:`Сохранение состояния:`}),` например, в обработчиках событий или колбэках.`]})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Практический пример: Фабрика функций`}),(0,n.jsx)(e,{code:`function createMultiplier(multiplier) {
  return function(num) {
    return num * multiplier;
  };
}

const double = createMultiplier(2);
const triple = createMultiplier(3);

console.log(double(5)); // 10
console.log(triple(5)); // 15`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Важно помнить`}),(0,n.jsx)(`p`,{children:`Замыкания потребляют память, так как переменные, на которые они ссылаются, не могут быть удалены сборщиком мусора, пока существует само замыкание.`})]})}export{r as default};