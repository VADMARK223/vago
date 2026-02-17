import{c as e,d as t}from"./index-5Iak_7Bb.js";import{t as n}from"./Book.module-DeRDuBeZ.js";import{t as r}from"./BookHashLink-BMl38gQc.js";var i=t();function a(){return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`nav`,{className:n.toc,children:(0,i.jsx)(r,{id:`why_not_enum`,children:`Почему не enum?`})}),(0,i.jsx)(e,{code:`export const ROLE = {
    user: 'user',
    moderator: 'moderator',
    admin: 'admin',
} as const

export type Role = typeof ROLE[keyof typeof ROLE];`}),(0,i.jsx)(`hr`,{}),(0,i.jsxs)(`h2`,{children:[`1. Что делает `,(0,i.jsx)(`code`,{children:`as const`})]}),(0,i.jsxs)(`p`,{children:[(0,i.jsx)(`code`,{children:`as const`}),` делает:`]}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[`Все поля `,(0,i.jsx)(`code`,{children:`readonly`})]}),(0,i.jsxs)(`li`,{children:[`Все значения становятся `,(0,i.jsx)(`b`,{children:`литеральными типами`})]})]}),(0,i.jsx)(e,{code:`{
  readonly user: "user"
  readonly moderator: "moderator"
  readonly admin: "admin"
}
`}),(0,i.jsxs)(`p`,{children:[`⚠️ Это важно:`,(0,i.jsx)(`br`,{}),`значение `,(0,i.jsx)(`code`,{children:`"user"`}),` теперь не просто `,(0,i.jsx)(`code`,{children:`string`}),`, а конкретный тип`,` `,(0,i.jsx)(`code`,{children:`"user"`}),`.`]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h2`,{children:`2. Создание типа`}),(0,i.jsx)(e,{code:`export type Role = typeof ROLE[keyof typeof ROLE];
`}),(0,i.jsxs)(`h3`,{children:[`Шаг1 - `,(0,i.jsx)(`code`,{children:`typeof ROLE`})]}),(0,i.jsx)(`p`,{children:`Это тип объекта:`}),(0,i.jsx)(e,{code:`{
  readonly user: "user"
  readonly moderator: "moderator"
  readonly admin: "admin"
}`}),(0,i.jsx)(`hr`,{}),(0,i.jsxs)(`h3`,{children:[`Шаг 2 — `,(0,i.jsx)(`code`,{children:`keyof typeof ROLE`})]}),(0,i.jsx)(`p`,{children:`keyof берёт ключи объекта:`}),(0,i.jsx)(e,{code:`"user" | "moderator" | "admin"`}),(0,i.jsx)(`hr`,{}),(0,i.jsxs)(`h3`,{children:[`Шаг 3 — `,(0,i.jsx)(`code`,{children:`typeof ROLE[keyof typeof ROLE]`})]}),(0,i.jsxs)(`p`,{children:[`Это называется `,(0,i.jsx)(`b`,{children:`indexed access type`}),`.`]}),(0,i.jsx)(`blockquote`,{children:`Возьми тип ROLE и получи типы всех значений по всем ключам.`}),(0,i.jsx)(e,{code:`"user" | "moderator" | "admin"`}),(0,i.jsx)(`hr`,{}),(0,i.jsxs)(`h3`,{children:[`🔥 В итоге тип `,(0,i.jsx)(`code`,{children:`Role`}),` равен:`]}),(0,i.jsx)(e,{code:`type Role = "user" | "moderator" | "admin"`}),(0,i.jsx)(`p`,{children:`И при этом:`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`не нужно вручную писать union`}),(0,i.jsx)(`li`,{children:`всё синхронизировано с объектом ROLE`})]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h2`,{id:`why_not_enum`,children:`🔥 Почему не enum?`}),(0,i.jsx)(`p`,{children:`Можно было бы так:`}),(0,i.jsx)(e,{code:`enum Role {
  User = "user",
  Moderator = "moderator",
  Admin = "admin"
}`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`не генерирует лишний JS код`}),(0,i.jsx)(`li`,{children:`проще`}),(0,i.jsx)(`li`,{children:`гибче`}),(0,i.jsx)(`li`,{children:`лучше работает с tree-shaking`}),(0,i.jsx)(`li`,{children:`удобнее в React / фронте`})]}),(0,i.jsx)(`hr`,{})]})}export{a as default};