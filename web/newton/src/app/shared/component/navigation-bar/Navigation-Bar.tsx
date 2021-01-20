import { AppNavBar, NavItemT, setItemActive } from 'baseui/app-nav-bar'
import { Overflow, DeleteAlt } from 'baseui/icon'
import { DefaultUser } from '../../../../internal/library/user/domain/user'
import { NavLink, useLocation } from 'react-router-dom'
import HorizontalLogo from '../horizontal-logo/Horizontal-Logo'
import React from 'react'

function NavigationBar() {
    const location = useLocation()
    const [mainItems, setMainItems] = React.useState<NavItemT[]>([
        { label: 'Home', info: { path: '/' } },
        { label: 'Explore', info: { path: '/explore' } }
    ])
    function handleMainItemSelect(item: NavItemT) {
        setMainItems(prev => setItemActive(prev, item));
    }
    return (
        <AppNavBar
            title={HorizontalLogo}
            mainItems={mainItems}
            onMainItemSelect={handleMainItemSelect}
            mapItemToNode={item => {
                item.active = item.info !== undefined ? location.pathname === item.info.path : item.active
                return (
                    <NavLink activeClassName='font-bold' exact to={item.info !== undefined ? item.info.path : '/'}>
                        {item.label}
                    </NavLink>
                )
            }}
            username={DefaultUser.DisplayName}
            usernameSubtitle={DefaultUser.PreferredUsername}
            userItems={[
                { icon: Overflow, label: 'Favorites' },
                { icon: DeleteAlt, label: 'Sign Out' }
            ]}
        // onUserItemSelect={item => console.log(item)}
        />
    )
}

export default NavigationBar