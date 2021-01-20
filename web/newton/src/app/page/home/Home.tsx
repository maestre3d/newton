import { Fragment } from 'react'
import { Helmet } from 'react-helmet'
import { APPLICATION_DOMAIN, APPLICATION_NAME } from '../../../internal/shared/domain/newton'

function Home() {
    return (
        <Fragment>
            <Helmet>
                <meta charSet='utf-8' />
                <title>{`${APPLICATION_NAME} - Home`}</title>
                <link rel='canonical' href={`https://${APPLICATION_DOMAIN}/`} />
            </Helmet>
            <div>Home</div>
        </Fragment>
    )
}

export default Home