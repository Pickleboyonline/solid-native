var banana = [0,1,3]
const adder = Deno.core.ops.op_adder
// Deno.core.print(Object.keys(Deno.core.ops).join('\n'));

Deno.core.print('' + adder(10))