import{c as e,d as t}from"./index-BWDKUK_q.js";var n=t();function r(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsxs)(`p`,{children:[(0,n.jsx)(`code`,{children:`flatMap`}),` = `,(0,n.jsx)(`b`,{children:`map`}),` + `,(0,n.jsx)(`b`,{children:`flatten`}),` (сплющивание)`]}),(0,n.jsxs)(`ol`,{children:[(0,n.jsxs)(`li`,{children:[`Пробегается по массиву (как `,(0,n.jsx)(`code`,{children:`map`}),`)`]}),(0,n.jsx)(`li`,{children:`Возвращает массив`}),(0,n.jsx)(`li`,{children:`Потом автоматически убирает один уровень вложенности`})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Примеры`}),(0,n.jsx)(`h3`,{children:`Без flatMap`}),(0,n.jsx)(e,{code:`const arr = [1, 2, 3]

const result = arr.map(n => [n, n * 10])

console.log(result)
// [
//   [1, 10],
//   [2, 20],
//   [3, 30]
// ]`}),(0,n.jsx)(`p`,{children:`Получился массив массивов 😬`}),(0,n.jsx)(`p`,{children:`Чтобы сделать его плоским:`}),(0,n.jsx)(e,{code:`const flat = result.flat()`}),(0,n.jsx)(`h3`,{children:`✅ С flatMap`}),(0,n.jsx)(e,{code:`const result = arr.flatMap(n => [n, n * 10])

console.log(result)
// [1, 10, 2, 20, 3, 30]`})]})}export{r as default};