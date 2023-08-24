(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[691],{8348:function(e,t,r){"use strict";r.d(t,{Z:function(){return s}});var n=r(5893);function o(e){var t=e.className;return(0,n.jsxs)("svg",{role:"status",className:"".concat(t," mr-2 animate-spin text-gray-400 fill-blue-600"),viewBox:"0 0 100 101",fill:"none",xmlns:"http://www.w3.org/2000/svg",children:[(0,n.jsx)("path",{d:"M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z",fill:"currentColor"}),(0,n.jsx)("path",{d:"M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z",fill:"currentFill"})]})}function s(){return(0,n.jsx)("div",{className:"flex items-center justify-center w-full h-full",children:(0,n.jsx)(o,{className:"w-12 h-12"})})}},1356:function(e,t,r){"use strict";r.r(t);var n=r(29),o=r(9499),s=r(7794),c=r.n(s),i=r(9734),l=r(8348),a=r(5893);function u(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter(function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable})),r.push.apply(r,n)}return r}function d(){return(0,a.jsx)("div",{className:"mx-auto cursor-not-allowed mt-24 bg-gray-100 shadow-lg border border-gray-200 rounded-lg w-64",children:(0,a.jsxs)("div",{className:"flex flex-col items-center justify-center p-3 text-center",children:[(0,a.jsx)("div",{children:"Collectors not deployed yet."}),(0,a.jsx)("div",{children:"Configure destinations to start"})]})})}function f(){return(f=(0,n.Z)(c().mark(function e(t){return c().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("/api/collectors",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({name:t})});case 2:if(e.sent.ok){e.next=5;break}throw Error("Failed to delete collector ".concat(t));case 5:case"end":return e.stop()}},e)}))).apply(this,arguments)}function h(e){var t=e.name,r=e.ready;return(0,a.jsx)("div",{className:"shadow-lg border border-gray-200 rounded-lg bg-white",children:(0,a.jsxs)("div",{className:"flex flex-col items-start p-5",children:[(0,a.jsx)("div",{className:"font-bold",children:t}),(0,a.jsxs)("div",{className:"flex flex-row justify-between w-full",children:[r?(0,a.jsx)("div",{className:"text-green-600 font-medium",children:"Ready"}):(0,a.jsx)("div",{className:"text-orange-400 font-medium",children:"Not Ready"}),(0,a.jsx)("button",{onClick:function(){return function(e){return f.apply(this,arguments)}(t)},className:"hover:bg-gray-100 cursor-pointer p-1 rounded-lg",children:(0,a.jsx)("svg",{xmlns:"http://www.w3.org/2000/svg",className:"h-5 w-5",viewBox:"0 0 20 20",fill:"currentColor",children:(0,a.jsx)("path",{fillRule:"evenodd",d:"M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z",clipRule:"evenodd"})})})]})]})})}t.default=function(){var e=(0,i.ZP)("/api/collectors",function(e){return fetch(e).then(function(e){return e.json()})}),t=e.data;return e.error?(0,a.jsx)("div",{children:"failed to load"}):t?(0,a.jsxs)("div",{className:"space-y-12",children:[(0,a.jsx)("div",{className:"text-4xl font-medium",children:"Active Collectors"}),t?(0,a.jsx)("div",{className:"grid lg:grid-cols-3 2xl:grid-cols-6 gap-4 pr-4",children:t.collectors.map(function(e){return(0,a.jsx)(h,function(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?u(Object(r),!0).forEach(function(t){(0,o.Z)(e,t,r[t])}):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):u(Object(r)).forEach(function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))})}return e}({},e),e.name)})}):(0,a.jsx)(d,{})]}):(0,a.jsx)(l.Z,{})}},6288:function(e,t,r){(window.__NEXT_P=window.__NEXT_P||[]).push(["/collectors",function(){return r(1356)}])},29:function(e,t,r){"use strict";function n(e,t,r,n,o,s,c){try{var i=e[s](c),l=i.value}catch(e){r(e);return}i.done?t(l):Promise.resolve(l).then(n,o)}function o(e){return function(){var t=this,r=arguments;return new Promise(function(o,s){var c=e.apply(t,r);function i(e){n(c,o,s,i,l,"next",e)}function l(e){n(c,o,s,i,l,"throw",e)}i(void 0)})}}r.d(t,{Z:function(){return o}})}},function(e){e.O(0,[734,774,888,179],function(){return e(e.s=6288)}),_N_E=e.O()}]);