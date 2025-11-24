#!/bin/bash

# Script to run all integration tests with interactive output, progress, and report generation

# Configuration
REPORT_DIR="test_reports"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
TEXT_REPORT="$REPORT_DIR/integration_test_report_${TIMESTAMP}.txt"
HTML_REPORT="$REPORT_DIR/integration_test_report_${TIMESTAMP}.html"
JSON_REPORT="$REPORT_DIR/integration_test_report_${TIMESTAMP}.json"
MD_REPORT="$REPORT_DIR/integration_test_report_${TIMESTAMP}.md"

# Create report directory
mkdir -p "$REPORT_DIR"

echo "Running all integration tests..."
echo "=================================="
echo ""

# Count total integration test files
TOTAL_TESTS=$(find . -name "*integration*test.go" -type f | wc -l)
echo "Found $TOTAL_TESTS integration test files"
echo "Reports will be saved to: $REPORT_DIR"
echo ""

# Capture full test output
TEST_OUTPUT=$(mktemp)
SUMMARY_FILE=$(mktemp)
trap "rm -f $TEST_OUTPUT $SUMMARY_FILE" EXIT

# Run tests and capture output with progress tracking
go test -v -count=1 ./... -run ".*Integration.*" 2>&1 | tee "$TEST_OUTPUT" | \
    awk -v total_tests="$TOTAL_TESTS" -v summary_file="$SUMMARY_FILE" '
    BEGIN { 
        passed=0; failed=0; total=0;
        passed_tests="";
        failed_tests="";
        print "Test Progress:\n"
    }
    /^=== RUN/ { 
        test_name=$0; 
        gsub(/^=== RUN   /, "", test_name);
        printf "[%d/%d] Running: %s\n", total+1, total_tests, test_name;
        total++;
        current_test=test_name;
    }
    /^--- PASS/ { 
        passed++; 
        passed_tests=passed_tests "\n  ✓ " current_test;
        printf "✓ PASSED (%d/%d - %.1f%%)\n\n", passed, total, (passed/total)*100;
    }
    /^--- FAIL/ { 
        failed++; 
        failed_tests=failed_tests "\n  ✗ " current_test;
        printf "✗ FAILED (%d/%d - %.1f%%)\n\n", failed, total, (failed/total)*100;
    }
    /^PASS$/ { 
        printf "Package PASSED\n\n";
    }
    /^FAIL$/ { 
        printf "Package FAILED\n\n";
    }
    END { 
        printf "\n==================================\n";
        printf "Summary:\n";
        printf "  Total: %d\n", total;
        printf "  Passed: %d (%.1f%%)\n", passed, (passed/total)*100;
        printf "  Failed: %d (%.1f%%)\n", failed, (failed/total)*100;
        printf "==================================\n";
        
        # Save summary
        print "PASSED_COUNT=" passed > summary_file
        print "FAILED_COUNT=" failed >> summary_file
        print "TOTAL_COUNT=" total >> summary_file
        print "PASSED_TESTS=\"" passed_tests "\"" >> summary_file
        print "FAILED_TESTS=\"" failed_tests "\"" >> summary_file
    }
    '

# Wait a moment for file to be written
sleep 0.5

# Read summary
if [ -f "$SUMMARY_FILE" ]; then
    source "$SUMMARY_FILE" 2>/dev/null
fi

# Set defaults if not set
PASSED_COUNT=${PASSED_COUNT:-0}
FAILED_COUNT=${FAILED_COUNT:-0}
TOTAL_COUNT=${TOTAL_COUNT:-0}
PASSED_TESTS=${PASSED_TESTS:-""}
FAILED_TESTS=${FAILED_TESTS:-""}

# Calculate percentages
if [ "$TOTAL_COUNT" -gt 0 ]; then
    PASS_PCT=$(awk "BEGIN {printf \"%.1f\", ($PASSED_COUNT/$TOTAL_COUNT)*100}")
    FAIL_PCT=$(awk "BEGIN {printf \"%.1f\", ($FAILED_COUNT/$TOTAL_COUNT)*100}")
else
    PASS_PCT=0
    FAIL_PCT=0
fi

# Generate Text Report
{
    echo "Integration Test Report"
    echo "======================"
    echo "Generated: $(date)"
    echo "Timestamp: $TIMESTAMP"
    echo ""
    echo "Summary"
    echo "-------"
    echo "Total Tests: $TOTAL_COUNT"
    echo "Passed: $PASSED_COUNT ($PASS_PCT%)"
    echo "Failed: $FAILED_COUNT ($FAIL_PCT%)"
    echo ""
    if [ -n "$PASSED_TESTS" ]; then
        echo "Passed Tests:$PASSED_TESTS"
        echo ""
    fi
    if [ -n "$FAILED_TESTS" ]; then
        echo "Failed Tests:$FAILED_TESTS"
        echo ""
    fi
    echo "Detailed Output"
    echo "---------------"
    cat "$TEST_OUTPUT"
} > "$TEXT_REPORT"

# Extract failure details for each test
extract_failure_details() {
    local test_output="$1"
    local test_name="$2"
    
    # Extract content between "--- FAIL: test_name" and next "---" or "PASS"/"FAIL" line
    awk -v test_name="$test_name" 'BEGIN { 
        capturing=0
        details=""
    }
    /^--- FAIL: / {
        if (index($0, test_name) > 0) {
            capturing=1
            next
        } else if (capturing) {
            capturing=0
        }
    }
    /^--- PASS:/ {
        if (capturing) {
            capturing=0
        }
    }
    /^PASS$/ || /^FAIL$/ {
        if (capturing) {
            capturing=0
        }
    }
    capturing {
        if ($0 !~ /^--- FAIL:/) {
            details=details $0 "\n"
        }
    }
    END {
        if (length(details) > 0) {
            print details
        } else {
            print "No detailed error information available."
        }
    }' "$test_output"
}

# Generate HTML Report - Build in parts
HTML_TEMP=$(mktemp)
cat > "$HTML_TEMP" <<EOF
<!DOCTYPE html>
<html>
<head>
    <title>Integration Test Report - $TIMESTAMP</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        h1 { color: #333; border-bottom: 3px solid #4CAF50; padding-bottom: 10px; }
        .summary { display: flex; gap: 20px; margin: 20px 0; flex-wrap: wrap; }
        .stat-box { flex: 1; min-width: 200px; padding: 20px; border-radius: 5px; text-align: center; }
        .stat-box.total { background: #e3f2fd; border: 2px solid #2196F3; }
        .stat-box.passed { background: #e8f5e9; border: 2px solid #4CAF50; }
        .stat-box.failed { background: #ffebee; border: 2px solid #f44336; }
        .stat-box h2 { margin: 0; font-size: 2.5em; }
        .stat-box p { margin: 5px 0 0 0; color: #666; }
        .section { margin: 30px 0; }
        .section h2 { color: #555; border-bottom: 2px solid #ddd; padding-bottom: 5px; }
        .test-list { background: #fafafa; padding: 15px; border-radius: 5px; margin: 10px 0; }
        .test-item { padding: 8px; margin: 5px 0; border-left: 4px solid; padding-left: 15px; font-family: monospace; }
        .test-item.passed { border-color: #4CAF50; background: #e8f5e9; }
        .test-item.failed { border-color: #f44336; background: #ffebee; }
        .test-header { display: flex; justify-content: space-between; align-items: center; cursor: pointer; }
        .test-header:hover { opacity: 0.8; }
        .test-name { flex: 1; }
        .dropdown-btn { 
            background: #f44336; 
            color: white; 
            border: none; 
            padding: 5px 15px; 
            border-radius: 3px; 
            cursor: pointer; 
            font-size: 0.9em;
            margin-left: 10px;
        }
        .dropdown-btn:hover { background: #d32f2f; }
        .dropdown-btn.expanded { background: #1976d2; }
        .dropdown-btn.expanded:hover { background: #1565c0; }
        .dropdown-content { 
            display: none; 
            margin-top: 10px; 
            padding: 15px; 
            background: #fff3e0; 
            border-left: 4px solid #ff9800; 
            border-radius: 4px;
            max-height: 500px;
            overflow-y: auto;
        }
        .dropdown-content.show { display: block; }
        .dropdown-content pre { 
            background: #263238; 
            color: #aed581; 
            padding: 15px; 
            border-radius: 5px; 
            overflow-x: auto; 
            font-size: 0.85em;
            margin: 0;
        }
        .details { margin-top: 30px; }
        .details pre { background: #263238; color: #aed581; padding: 15px; border-radius: 5px; overflow-x: auto; font-size: 0.9em; }
        .timestamp { color: #999; font-size: 0.9em; }
    </style>
    <script>
        function toggleDetails(testId) {
            const content = document.getElementById('details-' + testId);
            const btn = document.getElementById('btn-' + testId);
            if (content.classList.contains('show')) {
                content.classList.remove('show');
                btn.classList.remove('expanded');
                btn.textContent = 'Show Details';
            } else {
                content.classList.add('show');
                btn.classList.add('expanded');
                btn.textContent = 'Hide Details';
            }
        }
    </script>
</head>
<body>
    <div class="container">
        <h1>Integration Test Report</h1>
        <p class="timestamp">Generated: $(date) | Timestamp: $TIMESTAMP</p>
        
        <div class="summary">
            <div class="stat-box total">
                <h2>$TOTAL_COUNT</h2>
                <p>Total Tests</p>
            </div>
            <div class="stat-box passed">
                <h2>$PASSED_COUNT</h2>
                <p>Passed ($PASS_PCT%)</p>
            </div>
            <div class="stat-box failed">
                <h2>$FAILED_COUNT</h2>
                <p>Failed ($FAIL_PCT%)</p>
            </div>
        </div>
        
        <div class="section">
            <h2>Passed Tests ($PASSED_COUNT)</h2>
            <div class="test-list">
EOF

    if [ -n "$PASSED_TESTS" ]; then
        echo "$PASSED_TESTS" | grep -v '^$' | sed 's/^  [✓✗] /<div class="test-item passed">/;s/$/<\/div>/'
    else
        echo "<p>No tests passed.</p>"
    fi

    cat <<EOF
            </div>
        </div>
        
        <div class="section">
            <h2>Failed Tests ($FAILED_COUNT)</h2>
            <div class="test-list">
EOF

    # Generate failed tests HTML
    if [ -n "$FAILED_TESTS" ]; then
        test_counter=0
        while IFS= read -r line || [ -n "$line" ]; do
            [ -z "$line" ] && continue
            
            test_name=$(echo "$line" | sed 's/^  [✓✗] //')
            test_id="test_$test_counter"
            test_counter=$((test_counter + 1))
            
            failure_details=$(extract_failure_details "$TEST_OUTPUT" "$test_name")
            failure_details_escaped=$(echo "$failure_details" | sed 's/&/\&amp;/g; s/</\&lt;/g; s/>/\&gt;/g; s/"/\&quot;/g; s/'"'"'/\&#39;/g')
            
            cat >> "$HTML_TEMP" <<INNEREOF
            <div class="test-item failed">
                <div class="test-header" onclick="toggleDetails('$test_id')">
                    <span class="test-name">✗ $test_name</span>
                    <button id="btn-$test_id" class="dropdown-btn">Show Details</button>
                </div>
                <div id="details-$test_id" class="dropdown-content">
                    <pre>$failure_details_escaped</pre>
                </div>
            </div>
INNEREOF
        done <<< "$FAILED_TESTS"
    else
        echo "<p>No tests failed.</p>" >> "$HTML_TEMP"
    fi

    cat >> "$HTML_TEMP" <<EOF
            </div>
        </div>
        
        <div class="section details">
            <h2>Detailed Test Output</h2>
            <pre>$(cat "$TEST_OUTPUT" | sed 's/&/\&amp;/g; s/</\&lt;/g; s/>/\&gt;/g')</pre>
        </div>
    </div>
</body>
</html>
EOF

# Move temp file to final location
mv "$HTML_TEMP" "$HTML_REPORT"

# Generate JSON Report
{
    echo "{"
    echo "  \"timestamp\": \"$TIMESTAMP\","
    echo "  \"generated\": \"$(date -Iseconds 2>/dev/null || date)\","
    echo "  \"summary\": {"
    echo "    \"total\": $TOTAL_COUNT,"
    echo "    \"passed\": $PASSED_COUNT,"
    echo "    \"failed\": $FAILED_COUNT,"
    echo "    \"pass_percentage\": $PASS_PCT,"
    echo "    \"fail_percentage\": $FAIL_PCT"
    echo "  },"
    echo "  \"tests\": {"
    echo -n "    \"passed\": ["
    
    # Generate passed tests array
    if [ -n "$PASSED_TESTS" ]; then
        passed_list=$(echo "$PASSED_TESTS" | grep -v '^$' | sed 's/^  [✓✗] //' | sed 's/"/\\"/g' | sed 's/^/      "/;s/$/"/' | paste -sd, - | sed 's/^/      /')
        if [ -n "$passed_list" ]; then
            echo ""
            echo "$passed_list"
        fi
    fi
    
    echo "    ],"
    echo -n "    \"failed\": ["
    
    # Generate failed tests array
    if [ -n "$FAILED_TESTS" ]; then
        failed_list=$(echo "$FAILED_TESTS" | grep -v '^$' | sed 's/^  [✓✗] //' | sed 's/"/\\"/g' | sed 's/^/      "/;s/$/"/' | paste -sd, - | sed 's/^/      /')
        if [ -n "$failed_list" ]; then
            echo ""
            echo "$failed_list"
        fi
    fi
    
    echo "    ]"
    echo "  }"
    echo "}"
} > "$JSON_REPORT"

# Generate Markdown Report
MD_TEMP=$(mktemp)
cat > "$MD_TEMP" <<EOF
# Integration Test Report

**Generated:** $(date)  
**Timestamp:** $TIMESTAMP

---

## Summary

| Metric | Count | Percentage |
|--------|-------|------------|
| **Total Tests** | $TOTAL_COUNT | 100% |
| **Passed** | $PASSED_COUNT | $PASS_PCT% |
| **Failed** | $FAILED_COUNT | $FAIL_PCT% |

---

## Passed Tests ($PASSED_COUNT)

EOF

    if [ -n "$PASSED_TESTS" ]; then
        echo "$PASSED_TESTS" | grep -v '^$' | while IFS= read -r line; do
            test_name=$(echo "$line" | sed 's/^  [✓✗] //')
            echo "- ✅ $test_name"
        done
    else
        echo "*No tests passed.*"
    fi

    cat <<EOF

---

## Failed Tests ($FAILED_COUNT)

EOF

    # Generate failed tests Markdown
    if [ -n "$FAILED_TESTS" ]; then
        test_counter=0
        while IFS= read -r line || [ -n "$line" ]; do
            [ -z "$line" ] && continue
            
            test_name=$(echo "$line" | sed 's/^  [✓✗] //')
            test_counter=$((test_counter + 1))
            test_name_escaped=$(echo "$test_name" | sed 's/[`*_]/\\&/g')
            failure_details=$(extract_failure_details "$TEST_OUTPUT" "$test_name")
            
            cat >> "$MD_TEMP" <<INNEREOF
### $test_counter. ❌ $test_name_escaped

<details>
<summary>Click to view error details</summary>

\`\`\`
$failure_details
\`\`\`

</details>

---

INNEREOF
        done <<< "$FAILED_TESTS"
    else
        echo "*No tests failed.*" >> "$MD_TEMP"
    fi

    cat >> "$MD_TEMP" <<EOF

---

## Detailed Test Output

<details>
<summary>Click to view full test output</summary>

\`\`\`
$(cat "$TEST_OUTPUT")
\`\`\`

</details>

---

*Report generated by integration test runner*
EOF

# Move temp file to final location
mv "$MD_TEMP" "$MD_REPORT"

echo ""
echo "=================================="
echo "Reports Generated:"
echo "  Text:  $TEXT_REPORT"
echo "  HTML:  $HTML_REPORT"
echo "  JSON:  $JSON_REPORT"
echo "  Markdown:  $MD_REPORT"
echo "=================================="
echo ""

EXIT_CODE=${PIPESTATUS[0]}
exit $EXIT_CODE
