import{c as e,d as t}from"./index-BWDKUK_q.js";import{t as n}from"./Book.module-DeRDuBeZ.js";import{t as r}from"./BookHashLink-sMMxGoyN.js";var i=t();function a(){return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`nav`,{className:n.toc,children:(0,i.jsx)(r,{id:`rules`,children:`Правила`})}),(0,i.jsx)(`h2`,{id:`rules`,children:`Правила`}),(0,i.jsxs)(`h3`,{id:`rules`,children:[`Правило №1 — Props = `,(0,i.jsx)(`code`,{children:`interface`})]}),(0,i.jsxs)(`p`,{children:[`Всегда описываем props компонентов через `,(0,i.jsx)(`code`,{children:`interface`}),`.`]}),(0,i.jsx)(e,{code:`interface CodeBlockProps {
  code: string
  lang?: CodeLang
}

export function CodeBlock({ code, lang = 'tsx' }: CodeBlockProps) {}`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{id:`rules`,children:`Правило №2 — Все unions и enum-подобные штуки = type`}),(0,i.jsx)(e,{code:`type CodeLang = 'tsx' | 'ts' | 'js' | 'go'
type Theme = 'light' | 'dark'
type Size = 'sm' | 'md' | 'lg'`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{id:`rules`,children:`Правило №3 — DTO / API / утилитарные типы = type`}),(0,i.jsx)(e,{code:`type UserDto = {
  id: string
  email: string
}

type UpdateUserDto = Partial<UserDto>
type UserPreview = Pick<UserDto, 'id' | 'email'>`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{id:`rules`,children:`Правило №4 — Расширяемое → interface, закрытое → type`}),(0,i.jsx)(e,{code:`interface BaseButtonProps {
  disabled?: boolean
}

interface IconButtonProps extends BaseButtonProps {
  icon: ReactNode
}`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{id:`rules`,children:`Правило №5 — Не смешиваем без причины`}),(0,i.jsx)(`p`,{children:`❌ плохо:`}),(0,i.jsx)(e,{code:`type ButtonProps = { onClick: () => void }
interface ButtonExtra { size: 'sm' | 'md' }`}),(0,i.jsx)(`p`,{children:`✅ хорошо:`}),(0,i.jsx)(e,{code:`interface ButtonProps {
  onClick: () => void
  size: ButtonSize
}

type ButtonSize = 'sm' | 'md'`})]})}export{a as default};