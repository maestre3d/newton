import React from 'react'
import { ThemeContext, ThemeLabelAlt, ToggleTheme } from '../../../../internal/shared/infrastructure/theme'
import { AppNavBar, NavItemT, setItemActive } from 'baseui/app-nav-bar'
import { Overflow, DeleteAlt } from 'baseui/icon'
import { DefaultUser } from '../../../../internal/library/user/domain/user'
import { NavLink, useLocation } from 'react-router-dom'
import HorizontalLogo from '../horizontal-logo/Horizontal-Logo'

let setMainItems: React.Dispatch<React.SetStateAction<NavItemT[]>>

function handleItems(item: NavItemT) {
    setMainItems(prev => setItemActive(prev, item));
}

function NavigationBar() {
    const location = useLocation()
    const { theme, setCurrentTheme } = React.useContext(ThemeContext)
    const [mainItems, setMainItemsLocal] = React.useState<NavItemT[]>([
        { label: 'Home', info: { path: '/' } },
        { label: 'Explore', info: { path: '/explore' } }
    ])
    setMainItems = setMainItemsLocal

    return (
        <AppNavBar
            title={HorizontalLogo}
            mainItems={mainItems}
            onMainItemSelect={handleItems}
            mapItemToNode={item => {
                item.active = item.info !== undefined ? location.pathname === item.info.path : true
                return (
                    <NavLink activeClassName='font-bold' exact to={item.info !== undefined && item.info.path !== undefined ?
                        item.info.path : '/'}>
                        {item.label}
                    </NavLink>
                )
            }}
            username={DefaultUser.DisplayName}
            usernameSubtitle={DefaultUser.PreferredUsername}
            userItems={[
                { icon: Overflow, label: 'Favorites' },
                { icon: Overflow, label: ThemeLabelAlt(theme), info: { theme: true } },
                { icon: DeleteAlt, label: 'Sign Out' }
            ]}
            onUserItemSelect={item => {
                if (item.info !== undefined && item.info.theme) {
                    setCurrentTheme(ToggleTheme(theme))
                }
            }}
        />
    )
}

export default NavigationBar