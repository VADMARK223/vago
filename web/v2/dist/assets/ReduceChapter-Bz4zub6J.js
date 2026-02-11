import{d as e,n as t}from"./index-BvNZr30m.js";var n=e();function r(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsxs)(`p`,{children:[(0,n.jsx)(`code`,{children:`reduce`}),` - это метод массива, который превращает его в`,` `,(0,n.jsx)(`b`,{children:`одно итоговое значение`})]}),(0,n.jsx)(`p`,{children:`Это может быть:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsx)(`li`,{children:`число`}),(0,n.jsx)(`li`,{children:`строка`}),(0,n.jsx)(`li`,{children:`объект`}),(0,n.jsx)(`li`,{children:`другой массив`}),(0,n.jsx)(`li`,{children:`вообще что угодно`})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Синтаксис`}),(0,n.jsx)(t,{code:`array.reduce((accumulator, currentValue) => {
    return newAccumulator
}, initialValue)`}),(0,n.jsx)(`p`,{children:`Где:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`accumulator`}),` — накопленное значение`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`currentValue`}),` — текущий элемент массива`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`initialValue`}),` — стартовое значение`]})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Пример (простой)`}),(0,n.jsx)(`p`,{children:`Сумма чисел`}),(0,n.jsx)(t,{code:`const numbers = [1, 2, 3, 4]

const sum = numbers.reduce((acc, num) => {
    return acc + num
}, 0)

console.log(sum) // 10`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`Пример (сложный)`}),(0,n.jsx)(`p`,{children:`Группировка по ключам`}),(0,n.jsx)(t,{code:`type ChapterType = 'react' | 'ts' | 'js'

const grouped = chapters.reduce<Record<ChapterType, Chapter[]>>(
    (acc, chapter) => {
        acc[chapter.type].push(chapter)
        return acc
    },
    { react: [], ts: [], js: [] }
)`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`initialValue очень важен`}),(0,n.jsx)(`p`,{children:`Если не передать его:`}),(0,n.jsx)(t,{code:`numbers.reduce((acc, num) => acc + num)
`}),(0,n.jsx)(`p`,{children:`Тогда:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsx)(`li`,{children:`первый элемент станет acc`}),(0,n.jsx)(`li`,{children:`reduce начнётся со второго`})]}),(0,n.jsx)(`p`,{children:`Это может привести к багам.`})]})}export{r as default};