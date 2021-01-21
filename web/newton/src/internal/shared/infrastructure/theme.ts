import { strict } from 'assert'
import React from 'react'

export const THEME = {
    light: 'light',
    dark: 'dark'
}

export interface themeContextType {
    theme: string
    setCurrentTheme: (currentTheme: string) => void
}

const ThemeDefault: themeContextType = {
    theme: THEME.light,
    setCurrentTheme: () => {}
}

export const ThemeContext = React.createContext<themeContextType>(ThemeDefault)

export function ThemeLabelAlt(theme: string): string {
    return theme === THEME.light ? 'Dark Mode' : 'Light Mode'
}

export function ToggleTheme(theme: string): string {
    return theme === THEME.light ? THEME.dark : THEME.light
}

export const useTheme = (): themeContextType => {
    const [theme, setTheme] = React.useState(THEME.light)
    const setCurrentTheme = React.useCallback((currentTheme: string): void => {
        setTheme(currentTheme)
    }, [])
    
    return {
        theme,
        setCurrentTheme
    }
}