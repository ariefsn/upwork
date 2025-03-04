import i18n, { type Config } from 'sveltekit-i18n';

type Params = {
	faq?: string
	faqCount?: number
	faqSourceCodeLink?: string
	yearCount?: number
}

const locales = ['en', 'id'];

const loadJson = (key: string) => {
	return locales
		.map((locale) => ({ locale, key, loader: async () => (await import(`./i18n/${locale}/${key}.json`)).default }))
}

const config: Config<Params> = {
	initLocale: 'en',
	loaders: [
		...loadJson('app'),
		...loadJson('upwork'),
		...loadJson('faq'),
		...loadJson('user'),
	],
};

export const { t, locale, locales: i18nLocales, loading, loadTranslations } = new i18n(config);


export const changeLocale = (lang: string) => {
	locale.set(lang);
	localStorage.setItem('locale', lang);
}

export const getLocale = () => {
	const navigatorLang = window.navigator.language.split('-')
	let locale = localStorage.getItem('locale')
	if (navigatorLang.length > 1 && !locale) {
		locale = navigatorLang[0]
	}

	return localStorage.getItem('locale') || 'en';
}