import{d as e,t}from"./index-Db0aEsii.js";var n=e();function r(){return(0,n.jsxs)(`div`,{children:[(0,n.jsx)(`h2`,{id:`useMemo`,children:`useMemo`}),(0,n.jsxs)(`p`,{children:[`Хук`,` `,(0,n.jsx)(`b`,{children:`запоминает (мемоизирует) результат вычисления и пересчитывает его только тогда, когда меняются зависимости, чтобы React не пересчитывал дорогие значения на каждом рендере.`})]}),(0,n.jsxs)(`blockquote`,{children:[`❌ не пересчитывай тяжёлую логику на каждый ререндер`,(0,n.jsx)(`br`,{}),`✅ пересчитывай её только когда реально нужно`]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Синтаксис`}),(0,n.jsx)(t,{code:`const value = useMemo(() => {
  return expensiveCalculation(a, b)
}, [a, b])
`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[`функция внутри `,(0,n.jsx)(`code`,{children:`useMemo`}),` возвращает значение`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`[a, b]`}),` — зависимости`]}),(0,n.jsxs)(`li`,{children:[`если `,(0,n.jsx)(`code`,{children:`a`}),` и `,(0,n.jsx)(`code`,{children:`b`}),` не изменились → React вернёт старый результат`]})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Назначение`}),(0,n.jsx)(`h4`,{children:`1. Тяжелые вычисление`}),(0,n.jsx)(t,{code:`const sortedUsers = useMemo(() => {
  return users.sort((a, b) => a.age - b.age)
}, [users])
`}),(0,n.jsxs)(`p`,{children:[`Без `,(0,n.jsx)(`code`,{children:`useMemo`}),` сортировка будет выполняться на каждый ререндер, даже если`,` `,(0,n.jsx)(`code`,{children:`users`}),` те же самые.`]}),(0,n.jsx)(`hr`,{}),(0,n.jsxs)(`h4`,{children:[`2. Стабильные ссылки (часто с `,(0,n.jsx)(`code`,{children:`memo`}),`)`]}),(0,n.jsx)(t,{code:`const filteredTodos = useMemo(() => {
  return todos.filter(t => t.done)
}, [todos])
`}),(0,n.jsx)(`p`,{children:`Это важно, если:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[`ты передаёшь значение в `,(0,n.jsx)(`code`,{children:`React.memo`})]}),(0,n.jsxs)(`li`,{children:[`или в зависимости другого хука (`,(0,n.jsx)(`code`,{children:`useEffect`}),`, `,(0,n.jsx)(`code`,{children:`useCallback`}),`)`]})]}),(0,n.jsx)(`h4`,{children:`3. Избежать лишних ререндеров дочерних компонентов`}),(0,n.jsx)(t,{code:`const config = useMemo(() => ({
  theme: 'dark',
  pageSize: 20,
}), [])
`}),(0,n.jsxs)(`p`,{children:[`Без `,(0,n.jsx)(`code`,{children:`useMemo`}),` объект создаётся заново → дочерний компонент думает, что пропсы изменились.`]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Важные нюансы`}),(0,n.jsxs)(`p`,{children:[(0,n.jsx)(`code`,{children:`useMemo`}),` — `,(0,n.jsx)(`b`,{children:`не кеш навсегда`})]}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[`React `,(0,n.jsx)(`b`,{children:`может забыть значение`})]}),(0,n.jsx)(`li`,{children:`нельзя полагаться на него как на persistent cache`})]}),(0,n.jsxs)(`h4`,{children:[(0,n.jsx)(`code`,{children:`seMemo`}),` ≠ `,(0,n.jsx)(`code`,{children:`useCallback`})]}),(0,n.jsx)(t,{code:`useMemo(() => value, deps)      // запоминает значение
useCallback(() => fn, deps)     // запоминает функцию
`}),(0,n.jsx)(`p`,{children:`На самом деле:`}),(0,n.jsx)(t,{code:`useCallback(fn, deps)
// это то же самое, что
useMemo(() => fn, deps)
`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Типичная ошибка`}),(0,n.jsx)(t,{code:`useMemo(() => {
  doSomething()
}, [a])
`}),(0,n.jsxs)(`p`,{children:[`⚠️ `,(0,n.jsx)(`code`,{children:`useMemo`}),` `,(0,n.jsx)(`b`,{children:`должен возвращать значение`}),`, Если тебе нужен сайд-эффект → это`,` `,(0,n.jsx)(`code`,{children:`useEffect`})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Короткое правило`})]})}function i(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(`h2`,{id:`useEffect`,children:`useEffect`}),(0,n.jsx)(`p`,{children:`Это хук для побочных эффектов (side effects).`}),(0,n.jsx)(t,{code:`useEffect(() => {}, [])        // 1 раз (mount)
useEffect(() => {})            // каждый ререндер
useEffect(() => {}, [a, b])    // при изменении a или b`}),(0,n.jsx)(`p`,{children:`Частая ошибка`}),(0,n.jsx)(t,{code:`useEffect(() => {
  fetchData()
}, [])
`}),(0,n.jsxs)(`p`,{children:[`❌ А внутри `,(0,n.jsx)(`code`,{children:`fetchData`}),` используется `,(0,n.jsx)(`code`,{children:`props`}),` или `,(0,n.jsx)(`code`,{children:`state`})]}),(0,n.jsxs)(`p`,{children:[`👉 `,(0,n.jsx)(`b`,{children:`stale closure`}),` — эффект видит старые значения`]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Cleanup — must know`}),(0,n.jsx)(t,{code:`useEffect(() => {
  const id = setInterval(...)
  return () => clearInterval(id)
}, [])`}),(0,n.jsx)(`p`,{children:`Без cleanup:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsx)(`li`,{children:`утечки памяти`}),(0,n.jsx)(`li`,{children:`дублирующиеся подписки`}),(0,n.jsx)(`li`,{children:`баги «само по себе»`})]})]})}function a(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(`h2`,{id:`useState`,children:`useState`}),(0,n.jsxs)(`p`,{children:[`Хук для хранения и обновления `,(0,n.jsx)(`b`,{children:`локального состояния компонента.`})]})]})}function o(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(`h2`,{id:`useReducer`,children:`useReducer`}),(0,n.jsxs)(`p`,{children:[`Хук для управления `,(0,n.jsx)(`b`,{children:`более сложным состоянием`}),`, когда `,(0,n.jsx)(`code`,{children:`useState`}),` уже начинает путаться и перегружен.`]}),(0,n.jsxs)(`p`,{children:[`useReducer — это альтернатива `,(0,n.jsx)(`code`,{children:`useState`}),`, основанная на идее редьюсера (как в Redux):`]}),(0,n.jsx)(`blockquote`,{children:`есть state → action → reducer → новый state`}),(0,n.jsx)(`p`,{children:`Он особенно полезен, когда:`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[`состояние `,(0,n.jsx)(`b`,{children:`объект или вложенная структура`})]}),(0,n.jsx)(`li`,{children:`много действий, которые его меняют`}),(0,n.jsx)(`li`,{children:`логика обновления состояния не должна быть размазана по компоненту`}),(0,n.jsx)(`li`,{children:`хочется предсказуемости и читаемости`})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(t,{code:`const [state, dispatch] = useReducer(reducer, initialState)`}),(0,n.jsxs)(`ul`,{children:[(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`state`}),` — текущее состояние`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`dispatch(action)`}),` — отправка действия`]}),(0,n.jsxs)(`li`,{children:[(0,n.jsx)(`code`,{children:`reducer(state, action)`}),` — чистая функция, которая возвращает новый state`]})]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h3`,{children:`Простой пример`}),(0,n.jsx)(t,{code:`interface User {
    name: string
    surname: string
}

type Action =
    | { type: 'SET_NAME'; payload: string }
    | { type: 'SET_SURNAME'; payload: string }

const initialState: User = {
    name: 'Vadim',
    surname: 'Markitanov'
}

const reducer = (state: User, action: Action) => {
    switch (action.type) {
        case 'SET_NAME':
            return {...state, name: action.payload}
        case 'SET_SURNAME':
            return {...state, surname: action.payload}
        default:
            return state
    }
}

export function Page() {
    const [state, dispatch] = useReducer(reducer, initialState)

    return (
        <>
            <Button onClick={() => {
                dispatch({type: 'SET_NAME', payload: 'Oleg'})
            }}>Change name "{state.name}"</Button>
        </>
    )
}`})]})}function s(){return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(`a`,{href:`#useState`,children:`useState`}),(0,n.jsx)(`br`,{}),(0,n.jsx)(`a`,{href:`#useEffect`,children:`useEffect`}),(0,n.jsx)(`br`,{}),(0,n.jsx)(`a`,{href:`#useMemo`,children:`useMemo`}),(0,n.jsx)(`br`,{}),(0,n.jsx)(`a`,{href:`#useReducer`,children:`useReducer`}),(0,n.jsx)(`br`,{}),(0,n.jsx)(`a`,{href:`#useImperativeHandle`,children:`useImperativeHandle`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(a,{}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(i,{}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`useContext`}),(0,n.jsx)(`p`,{children:`доступ к контексту`}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`useRef`}),(0,n.jsxs)(`p`,{children:[`Позволяет `,(0,n.jsx)(`b`,{children:`хранить мутируемое значение между рендерами, не вызывая повторный рендер`}),`.`]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(r,{}),(0,n.jsx)(`h2`,{children:`useCallback`}),(0,n.jsxs)(`p`,{children:[`Это про `,(0,n.jsx)(`b`,{children:`мемоизацию функций`}),`, чтобы React не создавал их заново на каждом рендере. Часто идёт в паре с `,(0,n.jsx)(`code`,{children:`useMemo`}),` и `,(0,n.jsx)(`code`,{children:`React.memo`}),`.`]}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(o,{}),(0,n.jsx)(`hr`,{}),(0,n.jsx)(`h2`,{children:`useLayoutEffect`}),(0,n.jsx)(`p`,{children:`sync-эффекты`}),(0,n.jsx)(`h2`,{children:`useTransition`}),(0,n.jsx)(`p`,{children:`приоритеты UI`}),(0,n.jsx)(`h2`,{children:`useDeferredValue`}),(0,n.jsx)(`p`,{children:`отложенные обновления`}),(0,n.jsx)(`h2`,{children:`useId`}),(0,n.jsx)(`p`,{children:`Генерирует стабильные уникальные id.`}),(0,n.jsx)(`h2`,{id:`useImperativeHandle`,children:`useImperativeHandle`}),(0,n.jsx)(`p`,{children:`Позволяет явно управлять API ref’а, передаваемого родителю.`})]})}export{s as default};