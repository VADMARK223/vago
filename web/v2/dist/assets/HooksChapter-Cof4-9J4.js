import{m as e,s as t}from"./index-B7oVVJHg.js";import{t as n}from"./Book.module-DeRDuBeZ.js";import{t as r}from"./BookHashLink-Bm4i5FcO.js";var i=e();function a(){return(0,i.jsxs)(`div`,{children:[(0,i.jsx)(`h2`,{id:`useMemo`,children:`6. Хук useMemo`}),(0,i.jsxs)(`p`,{children:[`Хук`,` `,(0,i.jsx)(`b`,{children:`запоминает (мемоизирует) результат вычисления и пересчитывает его только тогда, когда меняются зависимости, чтобы React не пересчитывал дорогие значения на каждом рендере.`})]}),(0,i.jsxs)(`blockquote`,{children:[`❌ не пересчитывай тяжёлую логику на каждый ререндер`,(0,i.jsx)(`br`,{}),`✅ пересчитывай её только когда реально нужно`]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Синтаксис`}),(0,i.jsx)(t,{code:`const value = useMemo(() => {
  return expensiveCalculation(a, b)
}, [a, b])
`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[`функция внутри `,(0,i.jsx)(`code`,{children:`useMemo`}),` возвращает значение`]}),(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`[a, b]`}),` — зависимости`]}),(0,i.jsxs)(`li`,{children:[`если `,(0,i.jsx)(`code`,{children:`a`}),` и `,(0,i.jsx)(`code`,{children:`b`}),` не изменились → React вернёт старый результат`]})]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Назначение`}),(0,i.jsx)(`h4`,{children:`1. Тяжелые вычисление`}),(0,i.jsx)(t,{code:`const sortedUsers = useMemo(() => {
  return users.sort((a, b) => a.age - b.age)
}, [users])
`}),(0,i.jsxs)(`p`,{children:[`Без `,(0,i.jsx)(`code`,{children:`useMemo`}),` сортировка будет выполняться на каждый ререндер, даже если`,` `,(0,i.jsx)(`code`,{children:`users`}),` те же самые.`]}),(0,i.jsx)(`hr`,{}),(0,i.jsxs)(`h4`,{children:[`2. Стабильные ссылки (часто с `,(0,i.jsx)(`code`,{children:`memo`}),`)`]}),(0,i.jsx)(t,{code:`const filteredTodos = useMemo(() => {
  return todos.filter(t => t.done)
}, [todos])
`}),(0,i.jsx)(`p`,{children:`Это важно, если:`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[`ты передаёшь значение в `,(0,i.jsx)(`code`,{children:`React.memo`})]}),(0,i.jsxs)(`li`,{children:[`или в зависимости другого хука (`,(0,i.jsx)(`code`,{children:`useEffect`}),`, `,(0,i.jsx)(`code`,{children:`useCallback`}),`)`]})]}),(0,i.jsx)(`h4`,{children:`3. Избежать лишних ререндеров дочерних компонентов`}),(0,i.jsx)(t,{code:`const config = useMemo(() => ({
  theme: 'dark',
  pageSize: 20,
}), [])
`}),(0,i.jsxs)(`p`,{children:[`Без `,(0,i.jsx)(`code`,{children:`useMemo`}),` объект создаётся заново → дочерний компонент думает, что пропсы изменились.`]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Важные нюансы`}),(0,i.jsxs)(`p`,{children:[(0,i.jsx)(`code`,{children:`useMemo`}),` — `,(0,i.jsx)(`b`,{children:`не кеш навсегда`})]}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[`React `,(0,i.jsx)(`b`,{children:`может забыть значение`})]}),(0,i.jsx)(`li`,{children:`нельзя полагаться на него как на persistent cache`})]}),(0,i.jsxs)(`h4`,{children:[(0,i.jsx)(`code`,{children:`seMemo`}),` ≠ `,(0,i.jsx)(`code`,{children:`useCallback`})]}),(0,i.jsx)(t,{code:`useMemo(() => value, deps)      // запоминает значение
useCallback(() => fn, deps)     // запоминает функцию
`}),(0,i.jsx)(`p`,{children:`На самом деле:`}),(0,i.jsx)(t,{code:`useCallback(fn, deps)
// это то же самое, что
useMemo(() => fn, deps)
`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Типичная ошибка`}),(0,i.jsx)(t,{code:`useMemo(() => {
  doSomething()
}, [a])
`}),(0,i.jsxs)(`p`,{children:[`⚠️ `,(0,i.jsx)(`code`,{children:`useMemo`}),` `,(0,i.jsx)(`b`,{children:`должен возвращать значение`}),`, Если тебе нужен сайд-эффект → это`,` `,(0,i.jsx)(`code`,{children:`useEffect`})]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Короткое правило`})]})}function o(){return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`h2`,{id:`useEffect`,children:`2. Хук useEffect`}),(0,i.jsx)(`p`,{children:`Это хук для побочных эффектов (side effects).`}),(0,i.jsx)(t,{code:`useEffect(() => {}, [])        // 1 раз (mount)
useEffect(() => {})            // каждый ререндер
useEffect(() => {}, [a, b])    // при изменении a или b`}),(0,i.jsx)(`p`,{children:`Частая ошибка`}),(0,i.jsx)(t,{code:`useEffect(() => {
  fetchData()
}, [])
`}),(0,i.jsxs)(`p`,{children:[`❌ А внутри `,(0,i.jsx)(`code`,{children:`fetchData`}),` используется `,(0,i.jsx)(`code`,{children:`props`}),` или `,(0,i.jsx)(`code`,{children:`state`})]}),(0,i.jsxs)(`p`,{children:[`👉 `,(0,i.jsx)(`b`,{children:`stale closure`}),` — эффект видит старые значения`]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Cleanup - must know`}),(0,i.jsx)(t,{code:`useEffect(() => {
  const id = setInterval(...)
  return () => clearInterval(id)
}, [])`}),(0,i.jsx)(`p`,{children:`Без cleanup:`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`утечки памяти`}),(0,i.jsx)(`li`,{children:`дублирующиеся подписки`}),(0,i.jsx)(`li`,{children:`баги «само по себе»`})]})]})}function s(){return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`h2`,{id:`useState`,children:`1. Хук useState`}),(0,i.jsxs)(`p`,{children:[`Хук для хранения и обновления `,(0,i.jsx)(`b`,{children:`локального состояния компонента.`})]})]})}function c(){return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`h2`,{id:`useReducer`,children:`8. Хук useReducer`}),(0,i.jsxs)(`p`,{children:[`Хук для управления `,(0,i.jsx)(`b`,{children:`более сложным состоянием`}),`, когда `,(0,i.jsx)(`code`,{children:`useState`}),` уже начинает путаться и перегружен.`]}),(0,i.jsxs)(`p`,{children:[`useReducer — это альтернатива `,(0,i.jsx)(`code`,{children:`useState`}),`, основанная на идее редьюсера (как в Redux):`]}),(0,i.jsx)(`blockquote`,{children:`есть state → action → reducer → новый state`}),(0,i.jsx)(`p`,{children:`Он особенно полезен, когда:`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[`состояние `,(0,i.jsx)(`b`,{children:`объект или вложенная структура`})]}),(0,i.jsx)(`li`,{children:`много действий, которые его меняют`}),(0,i.jsx)(`li`,{children:`логика обновления состояния не должна быть размазана по компоненту`}),(0,i.jsx)(`li`,{children:`хочется предсказуемости и читаемости`})]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(t,{code:`const [state, dispatch] = useReducer(reducer, initialState)`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`state`}),` — текущее состояние`]}),(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`dispatch(action)`}),` — отправка действия`]}),(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`reducer(state, action)`}),` — чистая функция, которая возвращает новый state`]})]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h3`,{children:`Простой пример`}),(0,i.jsx)(t,{code:`interface User {
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
}`})]})}const l=()=>(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`h2`,{id:`useLayoutEffect`,children:`3. Хук useLayoutEffect`}),(0,i.jsx)(`p`,{children:`Это .`}),(0,i.jsx)(`h3`,{children:`🎯 Когда нужен useLayoutEffect?`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`измерить размер элемента`}),(0,i.jsx)(`li`,{children:`изменить стили до показа`}),(0,i.jsx)(`li`,{children:`предотвратить "мигание"`}),(0,i.jsx)(`li`,{children:`синхронно поправить layout`})]}),(0,i.jsx)(`h3`,{children:`⚠️ Почему нельзя всегда использовать useLayoutEffect?`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`блокирует отрисовку`}),(0,i.jsx)(`li`,{children:`синхронный`}),(0,i.jsx)(`li`,{children:`может ухудшить производительность`})]}),(0,i.jsxs)(`blockquote`,{children:[(0,i.jsx)(`code`,{children:`useLayoutEffect`}),` — это "перехватить момент ДО того, как браузер покажет кадр".`]})]}),u=()=>(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`h2`,{id:`useTansition`,children:`9. Хук useTransition`}),(0,i.jsxs)(`p`,{children:[`Это хук из React 18, который позволяет пометить обновление состояния как`,` `,(0,i.jsx)(`b`,{children:`неприоритетное (transition)`}),`. Ты контролируешь, какое обновление сделать низкоприоритетным`]}),(0,i.jsx)(`blockquote`,{children:`что-то должно обновиться, но не срочно — пусть React сделает это “когда сможет”, не блокируя интерфейс.`}),(0,i.jsx)(`h3`,{children:`Зачем он вообще нужен?`}),(0,i.jsxs)(`p`,{children:[`Без `,(0,i.jsx)(`code`,{children:`useTransition`}),`:`]}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`Пользователь печатает`}),(0,i.jsxs)(`li`,{children:[`Ты делаешь `,(0,i.jsx)(`code`,{children:`setState`})]}),(0,i.jsx)(`li`,{children:`Компонент рендерит 10 000 элементов`}),(0,i.jsx)(`li`,{children:`Input начинает лагать`})]}),(0,i.jsxs)(`p`,{children:[`С `,(0,i.jsx)(`code`,{children:`useTransition`}),`:`]}),(0,i.jsxs)(`ul`,{children:[(0,i.jsx)(`li`,{children:`Ввод остаётся плавным`}),(0,i.jsx)(`li`,{children:`Тяжёлый рендер выполняется “в фоне”`})]}),(0,i.jsx)(`h3`,{children:`Сигнатура`}),(0,i.jsx)(t,{code:`const [isPending, startTransition] = useTransition();`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`startTransition(fn)`}),` - оборачиваешь обновление`]}),(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`isPending`}),` — показывает, что transition ещё выполняется`]})]}),(0,i.jsx)(`h4`,{children:`Пример (поисковый фильтр)`}),(0,i.jsx)(t,{code:`const [query, setQuery] = useState('');
const [filtered, setFiltered] = useState(data);
const [isPending, startTransition] = useTransition();

const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  const value = e.target.value;
  setQuery(value); // срочное обновление

  startTransition(() => {
    setFiltered(
      data.filter(item => item.includes(value))
    );
  });
};`}),(0,i.jsx)(`p`,{children:`Теперь:`}),(0,i.jsxs)(`ul`,{children:[(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`query`}),` обновляется сразу (input не лагает)`]}),(0,i.jsxs)(`li`,{children:[(0,i.jsx)(`code`,{children:`filtered`}),` считается “в фоне”`]})]}),(0,i.jsx)(`p`,{children:`Можно показывать лоадер, пока transition не завершен`}),(0,i.jsx)(t,{code:`{isPending && <Spinner />}`})]}),d=()=>(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(`h2`,{id:`useDeferred`,children:`10. Хук useDeferred`}),(0,i.jsx)(`p`,{children:`Отложенные обновления. Ты говоришь: “вот это значение можно отложить”`})]});function f(){return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsxs)(`nav`,{className:n.toc,children:[(0,i.jsx)(r,{id:`useState`,children:`1. Хук useState`}),(0,i.jsx)(r,{id:`useEffect`,children:`2. Хук useEffect`}),(0,i.jsx)(r,{id:`useLayoutEffect`,children:`3. Хук useLayoutEffect`}),(0,i.jsx)(r,{id:`useContext`,children:`4. Хук useContext`}),(0,i.jsx)(r,{id:`useRef`,children:`5. Хук useRef`}),(0,i.jsx)(r,{id:`useMemo`,children:`6. Хук useMemo`}),(0,i.jsx)(r,{id:`useCallback`,children:` 7. Хук useCallback`}),(0,i.jsx)(r,{id:`useReducer`,children:`8. Хук useReducer`}),(0,i.jsx)(r,{id:`useTansition`,children:`9. Хук useTransition`}),(0,i.jsx)(r,{id:`useDeferred`,children:`10. Хук useDeferred`}),(0,i.jsx)(r,{id:`useId`,children:`11. Хук useId`}),(0,i.jsx)(r,{id:`useImperativeHandle`,children:`12. Хук useImperativeHandle`})]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(s,{}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(t,{code:`
📌 Один кадр браузера

JS task start
   ↓
React render (строим виртуальный DOM)
   ↓
React commit (обновляем реальный DOM + ref)
   ↓
useLayoutEffect  ← можно менять DOM ДО показа
   ↓
Browser layout (браузер считает размеры)
   ↓
Browser paint (показывает картинку пользователю)
   ↓
--------------------------- КАДР ПОКАЗАН ---------------------------
   ↓
useEffect        ← выполняется уже ПОСЛЕ того как пользователь увидел кадр

`}),(0,i.jsx)(o,{}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(l,{}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h2`,{id:`useContext`,children:`4. Хук useContext`}),(0,i.jsx)(`p`,{children:`доступ к контексту`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h2`,{id:`useRef`,children:`5. Хук useRef`}),(0,i.jsxs)(`p`,{children:[`Позволяет `,(0,i.jsx)(`b`,{children:`хранить mutable значение между рендерами, не вызывая повторный рендер`}),`.`]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(a,{}),(0,i.jsx)(`h2`,{id:`useCallback`,children:`7. Хук useCallback`}),(0,i.jsxs)(`p`,{children:[`Это про `,(0,i.jsx)(`b`,{children:`мемоизацию функций`}),`, чтобы React не создавал их заново на каждом рендере. Часто идёт в паре с `,(0,i.jsx)(`code`,{children:`useMemo`}),` и `,(0,i.jsx)(`code`,{children:`React.memo`}),`.`]}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(c,{}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(u,{}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(d,{}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h2`,{id:`useId`,children:`11. Хук useId`}),(0,i.jsx)(`p`,{children:`Генерирует стабильные уникальные id.`}),(0,i.jsx)(`hr`,{}),(0,i.jsx)(`h2`,{id:`useImperativeHandle`,children:`12. Хук useImperativeHandle`}),(0,i.jsx)(`p`,{children:`Позволяет явно управлять API ref’ а, передаваемого родителю.`})]})}export{f as default};