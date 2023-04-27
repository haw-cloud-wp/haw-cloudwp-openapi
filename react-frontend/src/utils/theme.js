
export function isDark() {
    return localStorage.theme === "dark";
}

export function updateTheme() {
    if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark')
    } else {
        document.documentElement.classList.remove('dark')
    }
}
export function toggleTheme() {
    localStorage.theme = localStorage.theme === "dark" ? "light" : "dark";
    updateTheme();
}