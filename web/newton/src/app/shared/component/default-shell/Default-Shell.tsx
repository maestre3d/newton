import * as React from 'react'
import { styled } from 'baseui'
import { Route, Switch } from 'react-router-dom'
import NavigationBar from '../navigation-bar/Navigation-Bar'

const Home = React.lazy(() => import('../../../page/home/Home'))
const Explore = React.lazy(() => import('../../../page/explore/Explore'))

const Shell = styled('div', {
    display: 'flex',
    alignItems: 'start',
})

function DefaultShell() {
    return (
        <Shell className='w-screen h-screen flex flex-col w-screen bg-white dark:bg-gray-700 dark:text-white'>
            <NavigationBar />
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