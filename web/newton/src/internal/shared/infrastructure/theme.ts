import React from 'react'

export const THEME = {
    light: 'light',
    dark: 'dark'
}

function getCachedTheme(): string {
    const theme = localStorage.theme
    if (theme !== THEME.dark && theme !== THEME.light) {
        return THEME.light
    }
    return theme
}

export function LoadThemeDOM() {
    const htmlRef = document.querySelector('html')
    if (htmlRef !== null) {
        if (localStorage.theme === THEME.dark || (!('theme' in localStorage) &&
            window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            localStorage.theme = THEME.dark
            htmlRef.classList.add(THEME.dark)
            return
        }
        htmlRef.classList.remove(THEME.dark)
    }
}

export interface themeContextType {
    theme: string
    setCurrentTheme: (currentTheme: string) => void
}

const ThemeDefault: themeContextType = {
    theme: getCachedTheme() || THEME.light,
    setCurrentTheme: () => { }
}

export const ThemeContext = React.createContext<themeContextType>(ThemeDefault)

export function ThemeLabelAlt(theme: string): string {
    return theme === THEME.light ? 'Dark Mode' : 'Light Mode'
}

export function ToggleTheme(theme: string): string {
    return theme === THEME.light ? THEME.dark : THEME.light
}

export const useTheme = (): themeContextType => {
    const [theme, setTheme] = React.useState(getCachedTheme())
    const setCurrentTheme = React.useCallback((currentTheme: string): void => {
        setTheme(currentTheme)
        localStorage.theme = currentTheme
        LoadThemeDOM()
    }, [])

    return {
        theme,
        setCurrentTheme
    }
}