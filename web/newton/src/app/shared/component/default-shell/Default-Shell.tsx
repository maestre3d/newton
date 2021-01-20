import * as React from 'react'
import { styled } from 'baseui'
import { AppNavBar, NavItemT, setItemActive } from 'baseui/app-nav-bar'
import { Overflow, DeleteAlt } from 'baseui/icon'
import { APPLICATION_NAME } from '../../../../internal/shared/domain/newton'
import { DefaultUser } from '../../../../internal/library/user/domain/user'
import { Link, NavLink, Route, Switch, useLocation } from 'react-router-dom'

const Home = React.lazy(() => import('../../../page/home/Home'))
const Explore = React.lazy(() => import('../../../page/explore/Explore'))

const Shell = styled('div', {
    display: 'flex',
    alignItems: 'start',
})

const Logo = (
    <Link to='/' className='flex flex-row items-center'>
        <span className='w-8 h-8 rounded-full bg-gray-300 mr-2' />
        <span className='font-bold'>{APPLICATION_NAME}</span>
    </Link>
)

function Navbar() {
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
            title={Logo}
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

function DefaultShell() {
    return (
        <Shell className='w-screen h-screen flex flex-col w-screen bg-white dark:bg-gray-700 dark:text-white'>
            <Navbar />
            <React.Suspense fallback={<div>Loading...</div>}>
                <Switch>
                    <Route path='/explore'>
                        <Explore />
                    </Route>
                    <Route path='/'>
                        <Home />
                    </Route>
                </Switch>
            </React.Suspense>
        </Shell>
    )
}

export default DefaultShell