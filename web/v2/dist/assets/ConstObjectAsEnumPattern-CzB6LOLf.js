import{t as e}from"./CodeBlock-BK_7U7Vx.js";import{u as t}from"./index-Dgjj6vUq.js";var n=t();function r(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(`a`,{href:`#1`,children:`Почему не enum?`}),(0,n.jsx)(`br`,{}),(0,n.jsx)(e,{code:`export const ROLE = {
    user: 'user',
    moderator: 'moderator',
    admin: 'admin',
} as const

export type Role = typeof ROLE[keyof typeof ROLE];`}),(0,n.jsx)(`hr`,{}),(0,n.jsxs)(`h2`,{children:[`1. Что делает `,(0,n.jsx)(`code`,{children:`as const`})]}),(0,n.jsxs)(`p`,{children:[(0,n.jsx)(`code`,{children:`as const`}),` делает:`]}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[`Все поля `,(0,n.jsx)(`code`,{children:`readonly`})]}),(0,n.jsxs)(`li`,{children:[`Все значения становятся `,(0,n.jsx)(`b`,{children:`литеральными типами`})]})]}),(0,n.jsx)(e,{code:`{
  readonly user: "user"
  readonly moderator: "moderator"
  readonly admin: "admin"
}
`}),(0,n.jsxs)(`p`,{children:[`⚠️ Это важно:`,(0,n.jsx)(`br`,{}),`значение `,(0,n.jsx)(`code`,{children:`"user"`}),` теперь не просто `,(0,n.jsx)(`code`,{children:`string`}),`, а конкретный тип `,(0,n.jsx)(`code`,{children:`"user"`}),`.`]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`2. Создание типа`}),(0,n.jsx)(e,{code:`export type Role = typeof ROLE[keyof typeof ROLE];
`}),(0,n.jsxs)(`h3`,{children:[`Шаг1 - `,(0,n.jsx)(`code`,{children:`typeof ROLE`})]}),(0,n.jsx)(`p`,{children:`Это тип объекта:`}),(0,n.jsx)(e,{code:`{
  readonly user: "user"
  readonly moderator: "moderator"
  readonly admin: "admin"
}`}),(0,n.jsx)(`hr`,{}),(0,n.jsxs)(`h3`,{children:[`Шаг 2 — `,(0,n.jsx)(`code`,{children:`keyof typeof ROLE`})]}),(0,n.jsx)(`p`,{children:`keyof берёт ключи объекта:`}),(0,n.jsx)(e,{code:`"user" | "moderator" | "admin"`}),(0,n.jsx)(`hr`,{}),(0,n.jsxs)(`h3`,{children:[`Шаг 3 — `,(0,n.jsx)(`code`,{children:`typeof ROLE[keyof typeof ROLE]`})]}),(0,n.jsxs)(`p`,{children:[`Это называется `,(0,n.jsx)(`b`,{children:`indexed access type`}),`.`]}),(0,n.jsx)(`blockquote`,{children:`Возьми тип ROLE и получи типы всех значений по всем ключам.`}),(0,n.jsx)(e,{code:`"user" | "moderator" | "admin"`}),(0,n.jsx)(`hr`,{}),(0,n.jsxs)(`h3`,{children:[`🔥 В итоге тип `,(0,n.jsx)(`code`,{children:`Role`}),` равен:`]}),(0,n.jsx)(e,{code:`type Role = "user" | "moderator" | "admin"`}),(0,n.jsx)(`p`,{children:`И при этом:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsx)(`li`,{children:`не нужно вручную писать union`}),(0,n.jsx)(`li`,{children:`всё синхронизировано с объектом ROLE`})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{id:`1`,children:`🔥 Почему не enum?`}),(0,n.jsx)(`p`,{children:`Можно было бы так:`}),(0,n.jsx)(e,{code:`enum Role {
  User = "user",
  Moderator = "moderator",
  Admin = "admin"
}`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsx)(`li`,{children:`не генерирует лишний JS код`}),(0,n.jsx)(`li`,{children:`проще`}),(0,n.jsx)(`li`,{children:`гибче`}),(0,n.jsx)(`li`,{children:`лучше работает с tree-shaking`}),(0,n.jsx)(`li`,{children:`удобнее в React / фронте`})]}),(0,n.jsx)(`hr`,{})]})}export{r as default};