{% include navigation.html %}
{% raw %}
# Conditionals in gotemplate

```go
#! @{is_true_1} := true
#! @{is_true_2} := true
#! @{is_false_1} := false
#! @{is_false_2} := false
#! @{false_string} := "false"

#! @-if($is_true_1)
    TestTrue
#! @-end

#! @-if($is_false_1)
    TestFalse
#! @-end

#! @-if(and($is_true_1, $is_true_2))
    TestTrueAndTrue
#! @-end

#! @-if(and($is_true_1, $is_false_1))
    TestTrueAndFalse
#! @-end

#! @-if(or($is_true_1, $is_false_1))
    TestTrueOrFalse
#! @-end

#! @-if($false_string)
    FalseStringIsTrue
#! @-end
```

will give:

```go
    TestTrue
    TestTrueAndTrue
    TestTrueOrFalse
    FalseStringIsTrue
```

{% endraw %}