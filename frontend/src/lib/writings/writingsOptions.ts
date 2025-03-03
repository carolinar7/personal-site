import type { WritingOption } from "./writingsTypes";

const options: WritingOption[] = [
  // {
  //   id: "svelte-router-nested-routers",
  //   title: "svelte-router: nested routers",
  //   description: "My first writing and also as small description of how I finally figured out how to use nested routers in svelte-router for this page.",
  //   date: new Date("2025-03-02"),
  // },
]

export const writingOptions: WritingOption[] = options.sort((a, b) => b.date.getTime() - a.date.getTime());