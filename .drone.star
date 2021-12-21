def main(ctx):
    return dict(
        kind="pipeline",
        type="docker",
        name="default",
        trigger=dict(
            branch="master",
        ),
        steps=[
            go_test(),
        ],
    )

def go_test():
    return dict(
        name="go test",
        image="golang:1.18beta1-alpine3.14",
        commands=[
            "go test -cover ./...",
        ],
    )
