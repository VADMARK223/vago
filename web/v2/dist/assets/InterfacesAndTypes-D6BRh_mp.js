import{t as e}from"./CodeBlock-BK_7U7Vx.js";import{u as t}from"./index-Dgjj6vUq.js";var n=t();function r(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(`a`,{href:`#rules`,children:`Привила`}),(0,n.jsx)(`br`,{}),(0,n.jsx)(`h2`,{id:`rules`,children:`Правила`}),(0,n.jsxs)(`h3`,{id:`rules`,children:[`Правило №1 — Props = `,(0,n.jsx)(`code`,{children:`interface`})]}),(0,n.jsxs)(`p`,{children:[`Всегда описываем props компонентов через `,(0,n.jsx)(`code`,{children:`interface`}),`.`]}),(0,n.jsx)(e,{code:`interface CodeBlockProps {
  code: string
  lang?: CodeLang
}

export function CodeBlock({ code, lang = 'tsx' }: CodeBlockProps) {}`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{id:`rules`,children:`Правило №2 — Все unions и enum-подобные штуки = type`}),(0,n.jsx)(e,{code:`type CodeLang = 'tsx' | 'ts' | 'js' | 'go'
type Theme = 'light' | 'dark'
type Size = 'sm' | 'md' | 'lg'`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{id:`rules`,children:`Правило №3 — DTO / API / утилитарные типы = type`}),(0,n.jsx)(e,{code:`type UserDto = {
  id: string
  email: string
}

type UpdateUserDto = Partial<UserDto>
type UserPreview = Pick<UserDto, 'id' | 'email'>`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{id:`rules`,children:`Правило №4 — Расширяемое → interface, закрытое → type`}),(0,n.jsx)(e,{code:`interface BaseButtonProps {
  disabled?: boolean
}

interface IconButtonProps extends BaseButtonProps {
  icon: ReactNode
}`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{id:`rules`,children:`Правило №5 — Не смешиваем без причины`}),(0,n.jsx)(`p`,{children:`❌ плохо:`}),(0,n.jsx)(e,{code:`type ButtonProps = { onClick: () => void }
interface ButtonExtra { size: 'sm' | 'md' }`}),(0,n.jsx)(`p`,{children:`✅ хорошо:`}),(0,n.jsx)(e,{code:`interface ButtonProps {
  onClick: () => void
  size: ButtonSize
}

type ButtonSize = 'sm' | 'md'`})]})}export{r as default};