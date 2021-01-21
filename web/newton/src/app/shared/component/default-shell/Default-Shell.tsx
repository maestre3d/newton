import * as React from 'react'
import { Block } from 'baseui/block'
import { Route, Switch } from 'react-router-dom'
import NavigationBar from '../navigation-bar/Navigation-Bar'

const Home = React.lazy(() => import('../../../page/home/Home'))
const Explore = React.lazy(() => import('../../../page/explore/Explore'))

function DefaultShell() {
    return (
        <Block className='w-screen h-screen flex flex-col' backgroundColor='backgroundPrimary' color='primaryA'>
            <NavigationBar />
            <React.Suspense fallback={<div>Loading...</div>}>
                <Switch>
                    <Route path='/explore' component={Explore} />
                    <Route path='/' component={Home} />
                </Switch>
            </React.Suspense>
        </Block>
    )
}

export default DefaultShell