import lume from "lume/mod.ts";
import wiki from "wiki/mod.ts";

const site = lume({
  location: new URL("https://pickleboyonline.github.com/solid-native"),
});

site.use(wiki());

export default site;
