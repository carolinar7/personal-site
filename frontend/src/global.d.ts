// Documentation: https://mdsvex.pngwn.io/docs#install-it
declare module '*.md' {
	import type { SvelteComponent } from 'svelte'

	export default class Comp extends SvelteComponent{}

	export const metadata: Record<string, unknown>
}