import * as React from "react";
import { PaymentCard } from "baseui/payment-card";
import {LightTheme, BaseProvider, styled} from 'baseui';
import {StatefulInput} from 'baseui/input';
import {Button} from 'baseui/button';
import {Upload} from 'baseui/icon';

const Centered = styled('div', {
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  height: '100%',
});

export function Foo() {
    const [value, setValue] = React.useState("");
    return (
        <BaseProvider theme={LightTheme}>
            <Centered>
            <StatefulInput />
            <PaymentCard
                value={value}
                onChange={(e: any) => setValue(e.target.value)}
                clearOnEscape
                placeholder="Please enter your credit card number..."
                />
            <p>
                <Button endEnhancer={() => <Upload size={24} />}>
                    End Enhancer
                </Button>
                </p>
            </Centered>
        </BaseProvider>
    );
}