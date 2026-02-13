import{d as e,n as t}from"./index-DfVnzSvh.js";import{t as n}from"./BookHashLink-CaSJeSpy.js";var r=e();function i(){return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsxs)(`p`,{children:[`В TypeScript существуют специальные типы, которые используются в специфических ситуациях:`,` `,(0,r.jsx)(n,{id:`any`,children:(0,r.jsx)(`code`,{children:`any`})}),`,`,` `,(0,r.jsx)(n,{id:`unknown`,children:(0,r.jsx)(`code`,{children:`unknown`})}),`,`,` `,(0,r.jsx)(n,{id:`void`,children:(0,r.jsx)(`code`,{children:`void`})}),`и`,` `,(0,r.jsx)(n,{id:`never`,children:(0,r.jsx)(`code`,{children:`never`})}),`.`]}),(0,r.jsx)(`h2`,{id:`any`,children:(0,r.jsx)(`code`,{children:`any`})}),(0,r.jsxs)(`p`,{children:[`Тип `,(0,r.jsx)(`code`,{children:`any`}),` позволяет переменной принимать абсолютно любое значение. Это фактически отключает проверку типов для данной переменной.`]}),(0,r.jsx)(t,{code:`let value: any = 4;
value = "hello";
value = true;

value.foo(); // Ошибки не будет при компиляции, но может упасть в рантайме`}),(0,r.jsx)(`p`,{children:(0,r.jsx)(`i`,{children:`Использование any не рекомендуется, так как это лишает вас преимуществ TypeScript.`})}),(0,r.jsx)(`hr`,{}),(0,r.jsx)(`h2`,{id:`unknown`,children:(0,r.jsx)(`code`,{children:`unknown`})}),(0,r.jsxs)(`p`,{children:[`Тип `,(0,r.jsx)(`code`,{children:`unknown`}),` — это безопасный аналог `,(0,r.jsx)(`code`,{children:`any`}),`. Мы можем присвоить ему любое значение, но не можем вызывать методы или обращаться к свойствам, пока не уточним тип (type narrowing).`]}),(0,r.jsx)(t,{code:`let value: unknown = "hello";

// value.toUpperCase(); // Ошибка!

if (typeof value === "string") {
  console.log(value.toUpperCase()); // Теперь можно!
}`}),(0,r.jsx)(`hr`,{}),(0,r.jsx)(`h2`,{id:`void`,children:(0,r.jsx)(`code`,{children:`void`})}),(0,r.jsxs)(`p`,{children:[`Используется в основном как тип возвращаемого значения функций, которые ничего не возвращают (или возвращают `,(0,r.jsx)(`code`,{children:`undefined`}),`).`]}),(0,r.jsx)(t,{code:`function logMessage(message: string): void {
  console.log(message);
}`}),(0,r.jsx)(`hr`,{}),(0,r.jsx)(`h2`,{id:`never`,children:(0,r.jsx)(`code`,{children:`never`})}),(0,r.jsxs)(`p`,{children:[`Тип `,(0,r.jsx)(`code`,{children:`never`}),` представляет значения, которые никогда не возникнут. Обычно используется для функций, которые всегда выбрасывают ошибку или имеют бесконечный цикл.`]}),(0,r.jsx)(t,{code:`function throwError(message: string): never {
  throw new Error(message);
}

function infiniteLoop(): never {
  while (true) {}
}`}),(0,r.jsxs)(`p`,{children:[`Также `,(0,r.jsx)(`code`,{children:`never`}),` полезен для проверки исчерпываемости (exhaustiveness checking) в`,` `,(0,r.jsx)(`code`,{children:`switch`}),`.`]}),(0,r.jsx)(t,{code:`type Shape = 'circle' | 'square';

function getArea(shape: Shape) {
  switch (shape) {
    case 'circle': return 1;
    case 'square': return 2;
    default:
      const _exhaustiveCheck: never = shape;
      return _exhaustiveCheck;
  }
}`}),(0,r.jsx)(`hr`,{})]})}export{i as default};