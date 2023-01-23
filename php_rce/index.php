<?php
  // Exploit: aHR0cDovL2xvY2FsaG9zdC9hcGkucGhwP3BhdHRlcm49L2EvZSZyZXBsYWNlPXN5c3RlbSglMjdpZCUyNyk7JnN1YmplY3Q9YWJjCg==
  // Solution: VXNlIGxhdGVyIHZlcnNpb24gb2YgUEhQLiBEb24ndCB1c2UgdnVsbmVyYWJsZSBmdW5jdGlvbnMuIENvZGUgc2Nhbi4K

  echo "<br >Welcome to the admin panel<br >";
 
  if (isset($_GET['pattern']) && isset($_GET['replace']) && isset($_GET['subject'])) {
    $pattern = $_GET['pattern'];
    $replacement = $_GET['replace'];
    $subject = $_GET['subject'];

    echo "Original: ".$subject ."</br>";
    echo "Replaced: ".preg_replace($pattern, $replacement, $subject);
  } else {
    die();
  }
?>
